package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armStartIP = "startIP"
	armEndIP   = "endIP"
)

const (
	tfStartIP = "start_ip"
	tfEndIP   = "end_ip"
)

// RedisFirewallRuleConfig returns config for azurerm_redis_firewall_rule
func RedisFirewallRuleConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfStartIP:  fn.LookUpString(nil, params, convert.ToString(r.Properties, armStartIP)),
		tfEndIP:    fn.LookUpString(nil, params, convert.ToString(r.Properties, armEndIP)),
	}
}
