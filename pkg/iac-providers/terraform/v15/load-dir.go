

package tfv15

import (
	commons "github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
)

// LoadIacDir starts traversing from the given rootDir and traverses through
// all the descendant modules present to create an output list of all the
// resources present in rootDir and descendant modules
func (*TfV15) LoadIacDir(absRootDir string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {

	return commons.NewTerraformDirectoryLoader(absRootDir, "0.15.0", options).LoadIacDir()
}

// Name returns name of the provider
func (*TfV15) Name() string {
	return "terraform"
}
