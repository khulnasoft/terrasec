

package iacprovider

import (
	"reflect"

	tfplanv1 "github.com/khulnasoft/terrasec/pkg/iac-providers/tfplan/v1"
)

// tfplan specific constants
const (
	tfplan                  supportedIacType    = "tfplan"
	tfplanV1                supportedIacVersion = "v1"
	tfplanDefaultIacVersion                     = tfplanV1
)

// register tfplan as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(tfplan, tfplanV1, tfplanDefaultIacVersion, reflect.TypeOf(tfplanv1.TFPlan{}))
}
