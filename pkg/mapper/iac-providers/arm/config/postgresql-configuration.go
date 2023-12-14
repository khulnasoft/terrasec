

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armSource = "source"
	armValue  = "value"
)

const tfValue = "value"

// PostgreSQLConfigurationConfig returns config for azurerm_postgresql_configuration
func PostgreSQLConfigurationConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfTags:     r.Tags,
		tfName:     convert.ToString(r.Properties, armSource),
		tfValue:    convert.ToString(r.Properties, armValue),
	}
}
