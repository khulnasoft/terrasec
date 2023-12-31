package commons

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	hclConfigs "github.com/hashicorp/terraform/configs"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

// LoadIacFile parses the given terraform file from the given file path
func LoadIacFile(absFilePath, terraformVersion string) (allResourcesConfig output.AllResourceConfigs, err error) {

	// new terraform config parser
	parser := hclConfigs.NewParser(afero.NewOsFs())

	// load current iac file
	hclFile, diags := parser.LoadConfigFile(absFilePath)
	if hclFile == nil {
		errMessage := fmt.Sprintf("error occurred while loading config file '%s'. error:\n%v\n", absFilePath, getErrorMessagesFromDiagnostics(diags))
		zap.S().Debug(errMessage)
		return allResourcesConfig, fmt.Errorf(errMessage)
	}

	if diags.HasErrors() {
		errMessage := fmt.Sprintf("failed to load iac file '%s'. error:\n%v\n", absFilePath, getErrorMessagesFromDiagnostics(diags))
		zap.S().Debug(errMessage)
		return allResourcesConfig, fmt.Errorf(errMessage)
	}

	// initialize normalized output
	allResourcesConfig = make(map[string][]output.ResourceConfig)

	// traverse through all current's resources
	for _, managedResource := range hclFile.ManagedResources {

		// create output.ResourceConfig from hclConfigs.Resource
		resourceConfig, err := CreateResourceConfig(managedResource)
		if err != nil {
			return allResourcesConfig, fmt.Errorf("failed to create ResourceConfig")
		}

		resourceConfig.TerraformVersion = terraformVersion
		managedResource.Provider = ResolveProvider(managedResource, hclFile.RequiredProviders)
		resourceConfig.ProviderVersion = GetProviderVersion(hclFile, managedResource.Provider, terraformVersion)
		// set module name
		// module name for the file scan will always be root
		resourceConfig.ModuleName = "root"

		// extract file name from path
		resourceConfig.Source = getFileName(resourceConfig.Source)

		// append to normalized output
		if _, present := allResourcesConfig[resourceConfig.Type]; !present {
			allResourcesConfig[resourceConfig.Type] = []output.ResourceConfig{resourceConfig}
		} else {
			allResourcesConfig[resourceConfig.Type] = append(allResourcesConfig[resourceConfig.Type], resourceConfig)
		}
	}

	// successful
	return allResourcesConfig, nil
}

// getFileName return file name from the given file path
func getFileName(path string) string {
	_, file := filepath.Split(path)
	return file
}

// getErrorMessagesFromDiagnostics should be called when diags.HasErrors is true
func getErrorMessagesFromDiagnostics(diags hcl.Diagnostics) string {
	var errMsgs []string
	for _, v := range diags {
		errMsgs = append(errMsgs, v.Error())
	}
	return strings.Join(errMsgs, "\n")
}
