package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armDNSPrefix         = "dnsPrefix"
	armAgentPoolProfiles = "agentPoolProfiles"
	armPoolName          = "name"
	armNodeCount         = "count"
	armVMSize            = "vmSize"
	armAddonProfiles     = "addonProfiles"
	armNetworkProfile    = "networkProfile"
	armNetworkPlugin     = "networkPlugin"
	armNetworkPolicy     = "networkPolicy"
)

const (
	tfDNSPrefix       = "dns_prefix"
	tfDefaultNodePool = "default_node_pool"
	tfNodeCount       = "node_count"
	tfVMSize          = "vm_size"
	tfAddonProfile    = "addon_profile"
	tfConfig          = "config"
	tfNetworkProfile  = "network_profile"
	tfNetworkPlugin   = "network_plugin"
	tfNetworkPolicy   = "network_policy"
)

// KubernetesClusterConfig returns config for azurerm_kubernetes_cluster.
func KubernetesClusterConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation:  fn.LookUpString(nil, params, r.Location),
		tfName:      fn.LookUpString(nil, params, r.Name),
		tfTags:      r.Tags,
		tfDNSPrefix: fn.LookUpString(vars, params, convert.ToString(r.Properties, armDNSPrefix)),
	}

	poolProfiles := convert.ToSlice(r.Properties, armAgentPoolProfiles)
	dnp := make([]map[string]interface{}, 0)
	for _, p := range poolProfiles {
		profile := p.(map[string]interface{})
		newPool := map[string]interface{}{
			tfName:      fn.LookUpString(vars, params, convert.ToString(profile, armPoolName)),
			tfVMSize:    fn.LookUpString(vars, params, convert.ToString(profile, armVMSize)),
			tfNodeCount: fn.LookUpFloat64(vars, params, convert.ToString(profile, armNodeCount)),
		}
		dnp = append(dnp, newPool)
	}
	cf[tfDefaultNodePool] = dnp

	addonProfiles := convert.ToMap(r.Properties, armAddonProfiles)
	aps := make(map[string]interface{})
	for key, value := range addonProfiles {
		addon := value.(map[string]interface{})
		profile := map[string]interface{}{
			tfEnabled: convert.ToBool(addon, "enabled"),
		}

		if cfg, ok := addon["config"]; ok {
			profile[tfConfig] = cfg.(map[string]interface{})
		}

		if key == "kubeDashboard" {
			aps["kube_dashboard"] = profile
		}
		aps[key] = profile
	}
	cf[tfAddonProfile] = aps

	netProfile := convert.ToMap(r.Properties, armNetworkProfile)
	cf[tfNetworkProfile] = map[string]string{
		tfNetworkPlugin: fn.LookUpString(vars, params, convert.ToString(netProfile, armNetworkPlugin)),
		tfNetworkPolicy: fn.LookUpString(vars, params, convert.ToString(netProfile, armNetworkPolicy)),
	}
	return cf
}
