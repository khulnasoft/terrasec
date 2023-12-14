package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

var (
	testDataDir = "testdata"

	testRules = Rules{
		ScanRules: []string{"rule.1", "rule.2", "rule.3", "rule.4", "rule.5"},
		SkipRules: []string{"rule.1"},
	}

	testCategoryList = Category{List: []string{"category.1", "category.2"}}

	testNotifier = Notifier{
		NotifierType: "webhook",
		NotifierConfig: map[string]interface{}{
			"url": "testurl1",
		},
	}

	testK8sAdmControl = K8sAdmissionControl{
		Dashboard:      true,
		DeniedSeverity: highSeverity.Level,
		Categories:     testCategoryList.List,
		SaveRequests:   true,
	}

	highSeverity = Severity{Level: "high"}
)

func TestNewTerrasecConfigReader(t *testing.T) {

	testPolicy := Policy{
		RepoPath: "rego-subdir",
		BasePath: "custom-path",
		RepoURL:  "https://repository/url",
		Branch:   "branch-name",
	}

	type args struct {
		fileName string
	}

	tests := []struct {
		name          string
		args          args
		want          *TerrasecConfigReader
		wantErr       bool
		assertGetters bool
		Policy
		notifications map[string]Notifier
		Rules
	}{
		{
			name: "empty config file",
			args: args{
				fileName: "",
			},
			want: &TerrasecConfigReader{},
		},
		{
			name: "nonexistent config file",
			args: args{
				fileName: "test",
			},
			wantErr: true,
			want:    &TerrasecConfigReader{},
		},
		{
			name: "invalid config file format",
			args: args{
				fileName: "test.invalid",
			},
			wantErr: true,
			want:    &TerrasecConfigReader{},
		},
		{
			name: "invalid toml config file",
			args: args{
				fileName: filepath.Join(testDataDir, "invalid.toml"),
			},
			wantErr: true,
			want:    &TerrasecConfigReader{},
		},
		{
			name: "invalid yaml config file",
			args: args{
				fileName: filepath.Join(testDataDir, "invalid.toml"),
			},
			wantErr: true,
			want:    &TerrasecConfigReader{},
		},
		{
			name: "valid toml config file with partial fields",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config.toml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
				},
			},
		},
		{
			name: "valid toml config file with all fields",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-all-fields.toml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:               testRules,
					Severity:            highSeverity,
					Category:            testCategoryList,
					K8sAdmissionControl: testK8sAdmControl,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid toml config file with all fields and severity defined",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-severity.toml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Severity: highSeverity,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid toml config file with all fields and categories defined",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-category.toml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Category: testCategoryList,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-all-fields.yaml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:               testRules,
					Severity:            highSeverity,
					Category:            testCategoryList,
					K8sAdmissionControl: testK8sAdmControl,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields and severity defined",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-severity.yml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Severity: highSeverity,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
		{
			name: "valid yaml config file with all fields and categories defined",
			args: args{
				fileName: filepath.Join(testDataDir, "terrasec-config-category.yaml"),
			},
			want: &TerrasecConfigReader{
				config: TerrasecConfig{
					Policy: testPolicy,
					Notifications: map[string]Notifier{
						"webhook1": testNotifier,
					},
					Rules:    testRules,
					Category: testCategoryList,
				},
			},
			assertGetters: true,
			notifications: map[string]Notifier{
				"webhook1": testNotifier,
			},
			Policy: testPolicy,
			Rules:  testRules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTerrasecConfigReader(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTerrasecConfigReader() got error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerrasecConfigReader() = got %v, want %v", got, tt.want)
			}
			if tt.assertGetters {
				if !reflect.DeepEqual(got.getPolicyConfig(), tt.Policy) || !reflect.DeepEqual(got.getNotifications(), tt.notifications) || !reflect.DeepEqual(got.getRules(), tt.Rules) {
					t.Errorf("NewTerrasecConfigReader() = got config: %v, notifications: %v, rules: %v want config: %v, notifications: %v, rules: %v", got.getPolicyConfig(), got.getNotifications(), got.getRules(), tt.Policy, tt.notifications, tt.Rules)
				}
			}
		})
	}
}
