

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armSku              = "sku"
	armFamily           = "family"
	armCapacity         = "capacity"
	armEnableNonSSLPort = "enableNonSslPort"
)

const (
	tfEnableNonSSLPort = "enable_non_ssl_port"
	tfCapacity         = "capacity"
	tfFamily           = "family"
)

// RedisCacheConfig returns config for azurerm_redis_cache
func RedisCacheConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	if enabledNonSSLPort, ok := fn.LookUp(nil, params, convert.ToString(r.Properties, armEnableNonSSLPort)).(bool); ok {
		cf[tfEnableNonSSLPort] = enabledNonSSLPort
	}

	s := convert.ToMap(r.Properties, armSku)
	cf[tfSkuName] = fn.LookUpString(nil, params, convert.ToString(s, tfName))
	cf[tfFamily] = fn.LookUpString(nil, params, convert.ToString(s, armFamily))
	cf[tfCapacity] = fn.LookUpFloat64(nil, params, convert.ToString(s, armCapacity))

	return cf
}
