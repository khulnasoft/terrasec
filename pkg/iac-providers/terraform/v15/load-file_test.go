package tfv15

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons/test"
	"github.com/khulnasoft/terrasec/pkg/utils"
)

var testDataDir = "testdata"
var emptyTfFilePath = filepath.Join(testDataDir, "empty.tf")

func TestLoadIacFile(t *testing.T) {

	testErrorString1 := `error occurred while loading config file 'not-there'. error:
<nil>: Failed to read file; The file "not-there" could not be read.
`

	testErrorString2 := fmt.Sprintf(`failed to load iac file '%s'. error:
%s:1,21-2,1: Invalid block definition; A block definition must have block content delimited by "{" and "}", starting on the same line as the block header.
%s:1,1-5: Unsupported block type; Blocks of type "some" are not expected here.
`, emptyTfFilePath, emptyTfFilePath, emptyTfFilePath)

	table := []struct {
		name     string
		filePath string
		options  map[string]interface{}
		tfv15    TfV15
		want     output.AllResourceConfigs
		wantErr  error
	}{
		{
			name:     "invalid filepath",
			filePath: "not-there",
			tfv15:    TfV15{},
			wantErr:  fmt.Errorf(testErrorString1),
		},
		{
			name:     "empty config",
			filePath: filepath.Join(testDataDir, "testfile"),
			tfv15:    TfV15{},
		},
		{
			name:     "invalid config",
			filePath: filepath.Join(testDataDir, "empty.tf"),
			tfv15:    TfV15{},
			wantErr:  fmt.Errorf(testErrorString2),
		},
		{
			name:     "depends_on",
			filePath: filepath.Join(testDataDir, "depends_on", "main.tf"),
			tfv15:    TfV15{},
		},
		{
			name:     "count",
			filePath: filepath.Join(testDataDir, "count", "main.tf"),
			tfv15:    TfV15{},
		},
		{
			name:     "for_each",
			filePath: filepath.Join(testDataDir, "for_each", "main.tf"),
			tfv15:    TfV15{},
		},
		{
			name:     "required_providers",
			filePath: filepath.Join(testDataDir, "required-providers", "main.tf"),
			tfv15:    TfV15{},
		},
		{
			name:     "required_providers with configuration alias",
			filePath: filepath.Join(testDataDir, "required-providers", "configuration-alias", "main.tf"),
			tfv15:    TfV15{},
		},
		{
			name:     "provider with only alias",
			filePath: filepath.Join(testDataDir, "provider-with-only-alias", "main.tf"),
			tfv15:    TfV15{},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := tt.tfv15.LoadIacFile(tt.filePath, tt.options)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", gotErr, tt.wantErr)
			}
		})
	}

	table2 := []struct {
		name         string
		tfConfigFile string
		tfJSONFile   string
		options      map[string]interface{}
		tfv15        TfV15
		wantErr      error
	}{
		{
			name:         "config1",
			tfConfigFile: filepath.Join(testDataDir, "tfconfigs", "config1.tf"),
			tfJSONFile:   filepath.Join(testDataDir, "tfjson", "config1.json"),
			tfv15:        TfV15{},
			wantErr:      nil,
		},
		{
			name:         "dummyconfig",
			tfConfigFile: filepath.Join(testDataDir, "dummyconfig", "dummyconfig.tf"),
			tfJSONFile:   filepath.Join(testDataDir, "tfjson", "dummyconfig.json"),
			tfv15:        TfV15{},
			wantErr:      nil,
		},
	}

	for _, tt := range table2 {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.tfv15.LoadIacFile(tt.tfConfigFile, tt.options)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", gotErr, tt.wantErr)
			}
			var want output.AllResourceConfigs

			wantBytes, _ := os.ReadFile(tt.tfJSONFile)
			if utils.IsWindowsPlatform() {
				wantBytes = utils.ReplaceWinNewLineBytes(wantBytes)
			}

			err := json.Unmarshal(wantBytes, &want)
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
