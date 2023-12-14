

package config

import (
	"time"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

// KeyVaultSecretConfig returns config for azurerm_key_vault_secret
func KeyVaultSecretConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	a := convert.ToMap(r.Properties, armAttributes)
	if i := a["exp"]; i != nil {
		t := time.Unix(int64(i.(float64)), 0)
		cf[tfExpirationDate] = t.Format(time.RFC3339)
	}
	return cf
}
