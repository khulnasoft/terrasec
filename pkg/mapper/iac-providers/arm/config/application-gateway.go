package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armWafConfiguration = "webApplicationFirewallConfiguration"

const (
	tfWafConfiguration = "waf_configuration"
	tfEnabled          = "enabled"
)

// ApplicationGatewayConfig returns config for azurerm_application_gateway
func ApplicationGatewayConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfLocation: fn.LookUpString(nil, params, tfLocation),
		tfTags:     r.Tags,
	}

	w := convert.ToMap(r.Properties, armWafConfiguration)
	cf[tfWafConfiguration] = map[string]interface{}{
		tfEnabled: convert.ToBool(w, armEnabled),
	}
	return cf
}
