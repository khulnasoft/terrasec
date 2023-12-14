

package helmv3

import "github.com/hashicorp/go-multierror"

// HelmV3 struct implements the IacProvider interface
type HelmV3 struct {
	errIacLoadDirs *multierror.Error
}

type helmChartData map[string]interface{}

const (
	helmValuesFilename = "values.yaml"
	helmTemplateDir    = "templates"
	helmTestDir        = "tests"
)
