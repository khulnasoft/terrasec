

package tfv12

import (
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	commons "github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons"
	"go.uber.org/zap"
)

// LoadIacDir starts traversing from the given rootDir and traverses through
// all the descendant modules present to create an output list of all the
// resources present in rootDir and descendant modules
func (*TfV12) LoadIacDir(absRootDir string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {
	zap.S().Warn("There may be a few breaking changes while working with terraform v0.12 files. For further information, refer to https://github.com/khulnasoft/terrasec/releases/v1.3.0")
	return commons.NewTerraformDirectoryLoader(absRootDir, version, options).LoadIacDir()
}

// Name returns name of the provider
func (*TfV12) Name() string {
	return "terraform"
}
