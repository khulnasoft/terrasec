package commons

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons/test"
	"github.com/khulnasoft/terrasec/pkg/utils"
)

func TestLoadIacFile(t *testing.T) {
	type args struct {
		absFilePath      string
		terraformVersion string
	}
	tests := []struct {
		name       string
		args       args
		outputJSON string
		wantErr    bool
	}{
		{
			name: "file with no provider defined",
			args: args{
				absFilePath:      filepath.Join(testDataDir, "terraform_iac_files", "with_no_provider.tf"),
				terraformVersion: "0.15.0",
			},
			outputJSON: filepath.Join(testDataDir, "tfjson", "output_no_provider_defined.json"),
		},
		{
			name: "file with provider config",
			args: args{
				absFilePath:      filepath.Join(testDataDir, "terraform_iac_files", "with_provider_config.tf"),
				terraformVersion: "0.15.0",
			},
			outputJSON: filepath.Join(testDataDir, "tfjson", "output_with_provider_config.json"),
		},
		{
			name: "file with required provider",
			args: args{
				absFilePath:      filepath.Join(testDataDir, "terraform_iac_files", "with_required_provider.tf"),
				terraformVersion: "0.15.0",
			},
			outputJSON: filepath.Join(testDataDir, "tfjson", "output_with_required_provider.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadIacFile(tt.args.absFilePath, tt.args.terraformVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadIacFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var want output.AllResourceConfigs

			// Read the expected value and unmarshal into want
			contents, _ := os.ReadFile(tt.outputJSON)
			if utils.IsWindowsPlatform() {
				contents = utils.ReplaceWinNewLineBytes(contents)
			}

			err = json.Unmarshal(contents, &want)
			if err != nil {
				t.Errorf("unexpected error unmarshalling want: %v", err)
			}

			match, err := test.IdenticalAllResourceConfigs(got, want)
			if err != nil {
				t.Errorf("unexpected error checking result: %v", err)
			}
			if !match {
				g, _ := json.MarshalIndent(got, "", "  ")
				w, _ := json.MarshalIndent(want, "", "  ")
				t.Errorf("got '%v', want: '%v'", string(g), string(w))
			}
		})
	}
}
