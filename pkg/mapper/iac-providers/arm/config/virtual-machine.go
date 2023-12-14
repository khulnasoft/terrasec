

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const armNetworkInterfaces = "networkInterfaces"
const tfNetworkInterfaceIDs = "network_interface_ids"

// VirtualMachineConfig returns config for azurerm_virtual_machine
func VirtualMachineConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	profile := convert.ToMap(r.Properties, armNetworkProfile)
	if interfaces, ok := profile[armNetworkInterfaces].([]interface{}); ok {
		iFaceIDs := []string{}
		for _, fs := range interfaces {
			iFace := fs.(map[string]interface{})
			iFaceIDs = append(iFaceIDs, iFace["id"].(string))
		}
		cf[tfNetworkInterfaceIDs] = iFaceIDs
	}
	return cf
}
