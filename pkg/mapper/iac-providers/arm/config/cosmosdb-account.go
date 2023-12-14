

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armIPAddressOrRange = "ipAddressOrRange"
	armIPRules          = "ipRules"
)

const (
	tfIPRangeFilter = "ip_range_filter"
)

// CosmosDBAccountConfig returns config for azurerm_cosmosdb_account
func CosmosDBAccountConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	var v string
	ipr := convert.ToSlice(r.Properties, armIPRules)
	for _, s := range ipr {
		m := s.(map[string]interface{})
		v = convert.ToString(m, armIPAddressOrRange)
		break
	}
	if v != "" {
		cf[tfIPRangeFilter] = v
	}
	return cf
}
