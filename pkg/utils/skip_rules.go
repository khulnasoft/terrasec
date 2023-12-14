

package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"go.uber.org/zap"
)

const (
	// TerrasecSkip key used to detect rules for skipping violations
	TerrasecSkip = "runterrasec.io/skip"
	// TerrasecSkipRule key used to detect the rule to be skipped
	TerrasecSkipRule = "rule"
	// TerrasecSkipComment key used to detect comment skipping a give rule
	TerrasecSkipComment = "comment"
	// SkipRulesPrefix used to identify and trim the skipping rule patterns
	SkipRulesPrefix = "#ts:skip="
	// MetaDataIDRegex pattern to match Rego Metadata ID
	MetaDataIDRegex = `(AC_)(AWS|AZURE|GCP|K8S|GITHUB|DOCKER)[_]([\d]{4})`
	// MetaDataReferenceIDRegex pattern to match Rego Metadata ReferenceID
	MetaDataReferenceIDRegex = `(([ A-Za-z0-9]+[.-]{1}){2,5})([\d]+)`
	// SkipRuleCommentRegex used to detect comments in skipped rule
	SkipRuleCommentRegex = `([ \t]+.*){0,1}`
)

var (
	ruleIDRegex                    = fmt.Sprintf("(%s|%s)", MetaDataReferenceIDRegex, MetaDataIDRegex)
	ruleIDPattern                  = regexp.MustCompile(ruleIDRegex)
	skipRulesPattern               = regexp.MustCompile(fmt.Sprintf("(%s%s%s)", SkipRulesPrefix, ruleIDRegex, SkipRuleCommentRegex))
	infileInstructionNotPresentLog = "%s not present for resource: %s"
)

// GetSkipRules returns a list of rules to be skipped. The rules to be skipped
// can be set in terraform resource config with the following pattern:
// #ts:skip=AWS.S3Bucket.DS.High.1043
// #ts:skip=AWS.S3Bucket.DS.High.1044 reason to skip the rule
// each rule and its optional comment must be in a new line
func GetSkipRules(body string) []output.SkipRule {
	var skipRules []output.SkipRule

	// check if any rules comments are present in body
	if !skipRulesPattern.MatchString(body) {
		return skipRules
	}

	// extract all commented skip rules
	comments := skipRulesPattern.FindAllString(body, -1)

	// extract rule ids from comments
	for _, c := range comments {
		c = strings.TrimPrefix(c, SkipRulesPrefix)
		skipRule := getSkipRuleObject(c)
		if skipRule != nil {
			skipRules = append(skipRules, *skipRule)
		}
	}
	return skipRules
}

func getSkipRuleObject(s string) *output.SkipRule {
	if s == "" {
		return nil
	}

	var skipRule output.SkipRule
	comment := ruleIDPattern.Split(s, 2)[1]
	skipRule.Rule = ruleIDPattern.FindString(strings.TrimSpace(s))
	skipRule.Comment = strings.TrimSpace(comment)

	return &skipRule
}

// ReadSkipRulesFromMap returns a list of rules to be skipped. The rules to be skipped
// can be set in annotations for kubernetes manifests and Resource Metadata in AWS cft:
// k8s:
// metadata:
//
//	annotations:
//	  runterrasec.io/skip: |
//	    [{"rule": "accurics.kubernetes.IAM.109", "comment": "reason to skip the rule"}]
//
// cft:
// Resource:
//
//	myResource:
//	  Metadata:
//	    runterrasec.io/skip: |
//	      [{"rule": "AC_AWS_047", "comment": "reason to skip the rule"}]
//
// cft json:
//
//	"Resource":{
//	  "myResource":{
//	    "Metadata":{
//	       "runterrasec.io/skip": "[{\"rule\":\"AWS.CloudFormation.Medium.0603\"}]"
//	    }
//	  }
//	}
//
// each rule and its optional comment must be a string containing an json array like
// [{rule: ruleID, comment: reason for skipping}]
func ReadSkipRulesFromMap(skipRulesMap map[string]interface{}, resourceID string) []output.SkipRule {

	var skipRulesFromMap interface{}
	var ok bool
	if skipRulesFromMap, ok = skipRulesMap[TerrasecSkip]; !ok {
		zap.S().Debugf(infileInstructionNotPresentLog, TerrasecSkip, resourceID)
		return nil
	}

	if rules, ok := skipRulesFromMap.(string); ok {
		skipRules := make([]output.SkipRule, 0)
		err := json.Unmarshal([]byte(rules), &skipRules)
		if err != nil {
			zap.S().Debugf("json string %s cannot be unmarshalled to []output.SkipRules struct schema", rules)
			return nil
		}
		return skipRules
	}

	zap.S().Debugf("%s must be a string containing an json array like [{rule: ruleID, comment: reason for skipping}]", TerrasecSkip)
	return nil
}
