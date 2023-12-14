

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armAdminUserEnabled = "adminUserEnabled"

const (
	tfSku          = "sku"
	tfAdminEnabled = "admin_enabled"
)

// ContainerRegistryConfig returns config for azurerm_container_registry
func ContainerRegistryConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfSku:      fn.LookUpString(nil, params, r.SKU.Name),
	}

	if adminEnabled, ok := fn.LookUp(nil, params, convert.ToString(r.Properties, armAdminUserEnabled)).(bool); ok {
		cf[tfAdminEnabled] = adminEnabled
	}

	return cf
}
