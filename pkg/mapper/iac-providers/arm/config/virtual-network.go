

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armSubnets              = "subnets"
	armProperties           = "properties"
	armAddressPrefix        = "addressPrefix"
	armNetworkSecurityGroup = "networkSecurityGroup"
)

const (
	tfSubnet        = "subnet"
	tfAddressPrefix = "address_prefix"
	tfSecurityGroup = "security_group,omitempty"
)

// VirtualNetworkConfig returns config for azurerm_virtual_network
func VirtualNetworkConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	subs := convert.ToSlice(r.Properties, armSubnets)
	subnets := make([]map[string]string, 0)
	for _, ss := range subs {
		s := ss.(map[string]interface{})
		prop := convert.ToMap(s, armProperties)

		sub := map[string]string{
			tfName:          fn.LookUpString(vars, params, convert.ToString(s, tfName)),
			tfAddressPrefix: fn.LookUpString(vars, params, convert.ToString(prop, armAddressPrefix)),
		}

		if nsg := convert.ToMap(prop, armNetworkSecurityGroup); nsg != nil {
			if sg, ok := fn.LookUp(vars, params, nsg["id"].(string)).(string); ok {
				sub[tfSecurityGroup] = sg
			}
		}
		subnets = append(subnets, sub)
	}
	cf[tfSubnet] = subnets

	return cf
}
