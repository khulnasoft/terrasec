

package config

import (
	"time"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armAttributes = "attributes"
const tfExpirationDate = "expiration_date"

// KeyVaultKeyConfig returns config for azurerm_key_vault_key
func KeyVaultKeyConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	attr := convert.ToMap(r.Properties, armAttributes)
	if i := attr["exp"]; i != nil {
		t := time.Unix(int64(i.(float64)), 0)
		cf[tfExpirationDate] = t.Format(time.RFC3339)
	}
	return cf
}
