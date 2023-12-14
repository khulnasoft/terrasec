

package utils

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// GetAbsPolicyConfigPaths transforms the provided policy base path and repo path into absolute paths
func GetAbsPolicyConfigPaths(policyBasePath, policyRepoPath string) (string, string, error) {
	absolutePolicyBasePath, err := GetAbsPath(policyBasePath)
	if err != nil {
		return policyBasePath, policyRepoPath, errors.Errorf("invalid policy path `%s`, error : `%v`", policyBasePath, err)
	}

	absolutePolicyRepoPath, err := GetAbsPath(policyRepoPath)
	if err != nil {
		return policyRepoPath, policyBasePath, errors.Errorf("invalid repository path `%s`, error : `%v`", policyRepoPath, err)
	}

	if strings.HasPrefix(absolutePolicyRepoPath, absolutePolicyBasePath) {
		return absolutePolicyBasePath, absolutePolicyRepoPath, nil
	}

	zap.S().Debugf("absolute rego_subdir path, `%s`, does not fall under base repo path's `%s` directory structure", absolutePolicyRepoPath, absolutePolicyBasePath)
	zap.S().Debugf("appending rego_subdir path: `%s` to the policy base path: `%s`. checking ...", policyRepoPath, policyBasePath)

	absolutePolicyRepoPath = filepath.Join(absolutePolicyBasePath, policyRepoPath)
	return absolutePolicyBasePath, absolutePolicyRepoPath, nil
}

// CheckPolicyType checks if supplied policy type matches desired policy types
func CheckPolicyType(rulePolicyType string, desiredPolicyTypes []string) bool {
	normDesiredPolicyTypes := make(map[string]bool, len(desiredPolicyTypes))
	normRulePolicyType := EnsureUpperCaseTrimmed(rulePolicyType)

	for _, desiredPolicyType := range desiredPolicyTypes {
		desiredPolicyType = EnsureUpperCaseTrimmed(desiredPolicyType)
		normDesiredPolicyTypes[desiredPolicyType] = true
	}

	if _, ok := normDesiredPolicyTypes["ALL"]; ok {
		return true
	}

	_, ok := normDesiredPolicyTypes[normRulePolicyType]
	return ok
}
