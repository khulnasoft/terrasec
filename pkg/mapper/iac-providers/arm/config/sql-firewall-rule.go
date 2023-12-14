

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armStartIPAddress = "startIpAddress"
	armEndIPAddress   = "endIpAddress"
)

const (
	tfStartIPAddress = "start_ip_address"
	tfEndIPAddress   = "end_ip_address"
)

// SQLFirewallRuleConfig returns config for azurerm_sql_firewall_rule
func SQLFirewallRuleConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:       fn.LookUpString(nil, params, r.Location),
		tfName:           fn.LookUpString(nil, params, r.Name),
		tfTags:           r.Tags,
		tfStartIPAddress: fn.LookUpString(nil, params, convert.ToString(r.Properties, armStartIPAddress)),
		tfEndIPAddress:   fn.LookUpString(nil, params, convert.ToString(r.Properties, armEndIPAddress)),
	}
}
