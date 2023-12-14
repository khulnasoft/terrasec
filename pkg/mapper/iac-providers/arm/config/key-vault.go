

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armEnableSoftDelete = "enableSoftDelete"

const (
	tfTenantID          = "tenant_id"
	tfSoftDeleteEnabled = "soft_delete_enabled"
)

// KeyVaultConfig returns config for azurerm_key_vault
func KeyVaultConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:          fn.LookUpString(nil, params, r.Location),
		tfName:              fn.LookUpString(nil, params, r.Name),
		tfTags:              r.Tags,
		tfSoftDeleteEnabled: convert.ToBool(r.Properties, armEnableSoftDelete),
		tfTenantID:          convert.ToString(params, armTenantID),
	}
}
