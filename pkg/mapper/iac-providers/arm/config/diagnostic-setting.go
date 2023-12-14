

package config

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armStorageAccountID = "storageAccountId"
	armCategory         = "category"
	armLogs             = "logs"
)

const (
	tfTargetResourceID = "target_resource_id"
	tfStorageAccountID = "storage_account_id"
	tfLog              = "log"
	tfCategory         = "category"
	tfRetentionPolicy  = "retention_policy"
	tfDays             = "days"
)

// DiagnosticSettingConfig returns config for azurerm_monitor_diagnostic_setting
func DiagnosticSettingConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation:         fn.LookUpString(nil, params, r.Location),
		tfName:             fn.LookUpString(nil, params, r.Name),
		tfTags:             r.Tags,
		tfTargetResourceID: fn.LookUpString(vars, params, getTargetResourceID(r.DependsOn)),
		tfStorageAccountID: fn.LookUpString(vars, params, convert.ToString(r.Properties, armStorageAccountID)),
	}

	logs := convert.ToSlice(r.Properties, armLogs)
	if len(logs) > 0 {
		tfLogMap := make([]map[string]interface{}, 0)
		for _, lg := range logs {
			mp := lg.(map[string]interface{})
			policy := convert.ToMap(mp, armRetentionPolicy)

			l := map[string]interface{}{
				tfEnabled:  convert.ToBool(mp, armEnabled),
				tfCategory: convert.ToString(mp, armCategory),
			}

			isEnabled := convert.ToBool(policy, armEnabled)
			if isEnabled {
				l[tfRetentionPolicy] = map[string]interface{}{
					tfEnabled: isEnabled,
					tfDays:    fn.LookUpFloat64(vars, params, convert.ToString(policy, armDays)),
				}
			} else {
				l[tfRetentionPolicy] = map[string]interface{}{
					tfEnabled: isEnabled,
				}
			}
		}
		cf[tfLog] = tfLogMap
	}
	return cf
}

func getTargetResourceID(deps []string) string {
	for _, d := range deps {
		if strings.Contains(d, "vault") {
			return d
		}
	}
	return ""
}
