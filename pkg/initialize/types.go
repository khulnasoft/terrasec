

package initialize

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/khulnasoft/terrasec/pkg/policy"
)

type environmentPolicy struct {
	regoTemplate     string
	metadataFileName string
	resourceType     string
	policyMetadata   policy.RegoMetadata
}

func newPolicy(ruleMetadata environmentPolicyMetadata) (environmentPolicy, error) {
	var policy environmentPolicy
	var templateArgs map[string]interface{}

	policy.regoTemplate = "package accurics\n\n" + ruleMetadata.RuleTemplate
	policy.metadataFileName = ruleMetadata.RuleReferenceID + ".json"
	policy.resourceType = ruleMetadata.ResourceType

	policy.policyMetadata.Name = ruleMetadata.RuleName
	policy.policyMetadata.File = ruleMetadata.RegoName + ".rego"
	policy.policyMetadata.ResourceType = ruleMetadata.ResourceType
	policy.policyMetadata.Severity = ruleMetadata.Severity
	policy.policyMetadata.Description = ruleMetadata.RuleDisplayName
	policy.policyMetadata.ReferenceID = ruleMetadata.RuleReferenceID
	policy.policyMetadata.ID = ruleMetadata.RuleReferenceID
	policy.policyMetadata.Category = ruleMetadata.Category
	policy.policyMetadata.Version = ruleMetadata.Version

	templateString, ok := ruleMetadata.RuleArgument.(string)
	if !ok {
		return policy, fmt.Errorf("incorrect rule argument type, must be a string")
	}
	err := json.Unmarshal([]byte(templateString), &templateArgs)
	if err != nil {
		return policy, fmt.Errorf("error occurred while unmarshaling rule arguments into map[string]interface{}, error: '%w'", err)
	}
	policy.policyMetadata.TemplateArgs = templateArgs

	return policy, nil
}

func (p environmentPolicy) getType() string {
	provider := strings.ToLower(p.resourceType)

	if strings.HasPrefix(provider, "azure") {
		return "azure"
	}

	if strings.HasPrefix(provider, "google") {
		return "gcp"
	}

	if strings.HasPrefix(provider, "kubernetes") {
		return "k8s"
	}

	return strings.Split(provider, "_")[0]
}

type environmentPolicyMetadata struct {
	RuleName        string      `json:"ruleName"`
	RegoName        string      `json:"ruleTemplateName"`
	RuleArgument    interface{} `json:"ruleArgument"`
	Severity        string      `json:"severity"`
	RuleDisplayName string      `json:"ruleDisplayName"`
	Category        string      `json:"category"`
	RuleReferenceID string      `json:"ruleReferenceId"`
	Version         int         `json:"version"`
	RuleTemplate    string      `json:"ruleTemplate"`
	ResourceType    string      `json:"resourceType"`
}
