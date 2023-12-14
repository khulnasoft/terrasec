package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armAccess               = "access"
	armDirection            = "direction"
	armProtocol             = "protocol"
	armSourceAddressPrefix  = "sourceAddressPrefix"
	armSourcePortRange      = "sourcePortRange"
	armDestinationPortRange = "destinationPortRange"
)

const (
	tfAccess               = "access"
	tfDirection            = "direction"
	tfProtocol             = "protocol"
	tfSourceAddressPrefix  = "source_address_prefix"
	tfSourcePortRange      = "source_port_range"
	tfDestinationPortRange = "destination_port_range"
)

// NetworkSecurityRuleConfig returns config for azurerm_network_security_rule
func NetworkSecurityRuleConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:             fn.LookUpString(nil, params, r.Location),
		tfName:                 fn.LookUpString(nil, params, r.Name),
		tfTags:                 r.Tags,
		tfAccess:               convert.ToString(r.Properties, armAccess),
		tfDirection:            convert.ToBool(r.Properties, armDirection),
		tfProtocol:             convert.ToString(r.Properties, armProtocol),
		tfSourceAddressPrefix:  convert.ToString(r.Properties, armSourceAddressPrefix),
		tfSourcePortRange:      convert.ToString(r.Properties, armSourcePortRange),
		tfDestinationPortRange: convert.ToString(r.Properties, armDestinationPortRange),
	}
}
