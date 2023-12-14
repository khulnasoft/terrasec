

package tfv15

import (
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	commons "github.com/khulnasoft/terrasec/pkg/iac-providers/terraform/commons"
)

// LoadIacFile parses the given terraform file from the given file path
func (*TfV15) LoadIacFile(absFilePath string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {
	return commons.LoadIacFile(absFilePath, version)
}
