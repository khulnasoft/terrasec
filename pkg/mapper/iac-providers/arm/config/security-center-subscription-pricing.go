

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armPricingTier = "pricingTier"
const tfTier = "tier"

// SecurityCenterSubscriptionPricingConfig returns config for azurerm_security_center_subscription_pricing
func SecurityCenterSubscriptionPricingConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfTier:     fn.LookUpString(nil, params, convert.ToString(r.Properties, armPricingTier)),
	}
}
