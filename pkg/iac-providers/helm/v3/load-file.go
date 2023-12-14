

package helmv3

import (
	"fmt"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"go.uber.org/zap"
)

var (
	errLoadIacFileNotSupported = fmt.Errorf("load iac file is not supported for helm")
)

// LoadIacFile is not supported for helm. Only loading chart directories are supported
func (h *HelmV3) LoadIacFile(absRootPath string, options map[string]interface{}) (allResourcesConfig output.AllResourceConfigs, err error) {
	zap.S().Errorf("load iac file is not supported for helm")
	return make(map[string][]output.ResourceConfig), errLoadIacFileNotSupported
}
