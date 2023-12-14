package tfplan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
)

const jqQuery = `[.planned_values.root_module | .. | select(.type? != null and .address? != null and .mode? == "managed") | {id: .address?, type: .type?, name: .name?, config: .values?, source: ""}]`

var (
	errIncorrectFormatVersion = fmt.Errorf("terraform format version should be one of '%s'", strings.Join(getTfPlanFormatVersions(), ", "))
	errEmptyTerraformVersion  = fmt.Errorf("terraform version cannot be empty in tfplan json")
)

func getTfPlanFormatVersions() []string {
	return []string{"0.1", "0.2"}
}

// LoadIacFile parses the given tfplan file from the given file path
func (t *TFPlan) LoadIacFile(absFilePath string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {

	zap.S().Debug("processing tfplan file")

	// read tfplan json file
	tfjson, err := os.ReadFile(absFilePath)
	if err != nil {
		errMsg := fmt.Sprintf("failed to read tfplan JSON file. error: '%v'", err)
		zap.S().Debug(errMsg)
		return allResourcesConfig, fmt.Errorf(errMsg)
	}

	// validate if provide file is a valid tfplan file
	if err := t.isValidTFPlanJSON(tfjson); err != nil {
		return allResourcesConfig, fmt.Errorf("invalid terraform json file; error: '%v'", err)
	}

	// run jq query on tfplan json
	processed, err := utils.JQFilterWithQuery(jqQuery, tfjson)
	if err != nil {
		errMsg := fmt.Sprintf("failed to process tfplan JSON. error: '%v'", err)
		zap.S().Debug(errMsg)
		return allResourcesConfig, fmt.Errorf(errMsg)
	}

	// decode processed out into output.ResourceConfig
	var resourceConfigs []output.ResourceConfig
	if err := json.Unmarshal(processed, &resourceConfigs); err != nil {
		errMsg := fmt.Sprintf("failed to decode processed jq output. error: '%v'", err)
		zap.S().Debug(errMsg)
		return allResourcesConfig, fmt.Errorf(errMsg)
	}

	// create AllResourceConfigs from resourceConfigs
	allResourcesConfig = make(map[string][]output.ResourceConfig)
	for _, r := range resourceConfigs {
		if _, present := allResourcesConfig[r.Type]; !present {
			allResourcesConfig[r.Type] = []output.ResourceConfig{r}
		} else {
			allResourcesConfig[r.Type] = append(allResourcesConfig[r.Type], r)
		}
	}

	// return output
	return allResourcesConfig, nil
}

// isValidTFPlanJSON validates whether the provided file is a valid tf json file
func (t *TFPlan) isValidTFPlanJSON(tfjson []byte) error {

	// decode tfjson into map[string]interface{}
	if err := json.Unmarshal(tfjson, &t); err != nil {
		return fmt.Errorf("failed to decode tfplan json. error: '%v'", err)
	}

	// check format version
	if !isValidVersion(t.FormatVersion) {
		return errIncorrectFormatVersion
	}

	// check terraform version
	if t.TerraformVersion == "" {
		return errEmptyTerraformVersion
	}

	return nil
}

func isValidVersion(v string) bool {
	for _, x := range getTfPlanFormatVersions() {
		if x == v {
			return true
		}
	}
	return false
}
