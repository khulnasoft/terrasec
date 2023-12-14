package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armTargetResourceID                         = "targetResourceId"
	armStorageID                                = "storageId"
	armEnabled                                  = "enabled"
	armRetentionPolicy                          = "retentionPolicy"
	armDays                                     = "days"
	armFlowAnalyticsConfiguration               = "flowAnalyticsConfiguration"
	armNetworkWatcherFlowAnalyticsConfiguration = "networkWatcherFlowAnalyticsConfiguration"
	armWorkspaceID                              = "workspaceId"
	armWorkspaceRegion                          = "workspaceRegion"
	armWorkspaceResourceID                      = "workspaceResourceId"
	armTrafficAnalyticsInterval                 = "trafficAnalyticsInterval"
)

const (
	tfNetworkSecurityGroupID = "network_security_group_id"
	tfTrafficAnalytics       = "traffic_analytics"
	tfWorkspaceID            = "workspace_id,omitempty"
	tfWorkspaceRegion        = "workspace_region,omitempty"
	tfWorkspaceResourceID    = "workspace_resource_id,omitempty"
	tfIntervalInMinutes      = "interval_in_minutes,omitempty"
)

// NetworkWatcherFlowLogConfig returns config for azurerm_network_watcher_flow_log
func NetworkWatcherFlowLogConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation:               fn.LookUpString(nil, params, r.Location),
		tfName:                   fn.LookUpString(nil, params, r.Name),
		tfTags:                   r.Tags,
		tfNetworkSecurityGroupID: fn.LookUpString(vars, params, convert.ToString(r.Properties, armTargetResourceID)),
		tfStorageAccountID:       fn.LookUpString(vars, params, convert.ToString(r.Properties, armStorageID)),
		tfEnabled:                convert.ToBool(r.Properties, armEnabled),
	}

	policy := convert.ToMap(r.Properties, armRetentionPolicy)
	cf[tfRetentionPolicy] = map[string]interface{}{
		tfEnabled: convert.ToBool(policy, armEnabled),
		tfDays:    fn.LookUpFloat64(vars, params, convert.ToString(policy, armDays)),
	}

	flowConfig := convert.ToMap(r.Properties, armFlowAnalyticsConfiguration)
	if flowConfig != nil {
		networkConfig := convert.ToMap(flowConfig, armNetworkWatcherFlowAnalyticsConfiguration)
		cf[tfTrafficAnalytics] = map[string]interface{}{
			tfEnabled:             convert.ToBool(networkConfig, armEnabled),
			tfWorkspaceID:         fn.LookUpString(vars, params, armWorkspaceID),
			tfWorkspaceRegion:     fn.LookUpString(vars, params, armWorkspaceRegion),
			tfWorkspaceResourceID: fn.LookUpString(vars, params, armWorkspaceResourceID),
			tfIntervalInMinutes:   convert.ToFloat64(networkConfig, armTrafficAnalyticsInterval),
		}
	}
	return cf
}
