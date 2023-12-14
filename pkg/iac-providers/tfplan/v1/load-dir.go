

package tfplan

import (
	"fmt"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
)

var (
	errIacDirNotSupport = fmt.Errorf("tfplan should always be a file, not a directory. Please specify path to tfplan file  with '-f' option")
)

// LoadIacDir is not supported for tfplan IacType. Terraform plan should always
// be a file and not a directory
func (k *TFPlan) LoadIacDir(absRootDir string, options map[string]interface{}) (output.AllResourceConfigs, error) {
	return output.AllResourceConfigs{}, errIacDirNotSupport
}

// Name returns name of the provider
func (*TFPlan) Name() string {
	return "tfplan"
}
