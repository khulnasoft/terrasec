package config

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/khulnasoft/terrasec/pkg/utils"
)

func TestLoadGlobalConfig(t *testing.T) {
	testConfigFile := filepath.Join(testDataDir, "terrasec-config-all-fields.toml")
	absDefaultBasePolicyPath, absDefaultPolicyRepoPath, _ := utils.GetAbsPolicyConfigPaths(defaultBasePolicyPath, defaultPolicyRepoPath)
	absCustomPath, absRegoSubdirPath, _ := utils.GetAbsPolicyConfigPaths("custom-path", "rego-subdir")

	type args struct {
		configFile string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		policyBasePath string
		policyRepoPath string
		repoURL        string
		branchName     string
		scanRules      []string
		skipRules      []string
		severity       string
		categories     []string
		notifications  map[string]Notifier
		k8sAdmControl  K8sAdmissionControl
	}{
		{
			// no error expected
			name: "global config file not specified",
			args: args{
				configFile: "",
			},
			policyBasePath: absDefaultBasePolicyPath,
			policyRepoPath: absDefaultPolicyRepoPath,
			repoURL:        defaultPolicyRepoURL,
			branchName:     defaultPolicyBranch,
		},
		{
			name: "global config file specified but doesn't exist",
			args: args{
				configFile: "test.toml",
			},
			wantErr:        true,
			policyBasePath: defaultBasePolicyPath,
			policyRepoPath: defaultPolicyRepoPath,
			repoURL:        defaultPolicyRepoURL,
			branchName:     defaultPolicyBranch,
		},
		{
			name: "valid global config file specified",
			args: args{
				configFile: testConfigFile,
			},
			policyBasePath: absCustomPath,
			policyRepoPath: absRegoSubdirPath,
			repoURL:        "https://repository/url",
			branchName:     "branch-name",
			scanRules:      testRules.ScanRules,
			skipRules:      testRules.SkipRules,
			severity:       highSeverity.Level,
			categories:     testCategoryList.List,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			k8sAdmControl: testK8sAdmControl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadGlobalConfig(tt.args.configFile); (err != nil) != tt.wantErr {
				t.Errorf("LoadGlobalConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if GetPolicyBasePath() != tt.policyBasePath || GetPolicyRepoPath() != tt.policyRepoPath || GetPolicyRepoURL() != tt.repoURL || GetPolicyBranch() != tt.branchName {
				t.Errorf("LoadGlobalConfig() error = got BasePath: %v, RepoPath: %v, RepoURL: %v, BranchName: %v, want BasePath: %v, RepoPath: %v, RepoURL: %v, BranchName: %v", GetPolicyBasePath(), GetPolicyRepoPath(), GetPolicyRepoURL(), GetPolicyBranch(), tt.policyBasePath, tt.policyRepoPath, tt.repoURL, tt.branchName)
			}

			if !utils.IsSliceEqual(GetScanRules(), tt.scanRules) || !utils.IsSliceEqual(GetSkipRules(), tt.skipRules) || !utils.IsSliceEqual(GetCategoryList(), tt.categories) || GetSeverityLevel() != tt.severity {
				t.Errorf("LoadGlobalConfig() error = got scan rules: %v, skip rules: %v, categories: %v, severity: %v, want scan rules: %v, skip rules: %v, categories: %v, severity: %v", GetScanRules(), GetSkipRules(), GetCategoryList(), GetSeverityLevel(), tt.scanRules, tt.skipRules, tt.categories, tt.severity)
			}

			if !reflect.DeepEqual(GetNotifications(), tt.notifications) || !reflect.DeepEqual(GetK8sAdmissionControl(), tt.k8sAdmControl) {
				t.Errorf("LoadGlobalConfig() error = got notifications: %v, k8s admission control: %v, want notifications: %v, k8s admission control: %v", GetNotifications(), GetK8sAdmissionControl(), tt.notifications, tt.k8sAdmControl)
			}
		})
	}
}
