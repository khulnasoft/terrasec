

package tfv12

import (
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	commons "github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons"
	"go.uber.org/zap"
)

// LoadIacFile parses the given terraform file from the given file path
func (*TfV12) LoadIacFile(absFilePath string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {
	zap.S().Warn("There may be a few breaking changes while working with terraform v0.12 files. For further information, refer to https://github.com/khulnasoft/terrasec/releases/v1.3.0")
	return commons.LoadIacFile(absFilePath, version)
}
