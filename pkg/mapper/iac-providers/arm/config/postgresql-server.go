package config

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

// PostgreSQLServerConfig returns config for azurerm_postgresql_server
func PostgreSQLServerConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfSkuName:  fn.LookUpString(vars, params, r.SKU.Name),
		tfVersion:  fn.LookUpString(vars, params, convert.ToString(r.Properties, armVersion)),
	}

	if profile := convert.ToMap(r.Properties, armStorageProfile); profile != nil {
		status := fn.LookUpString(vars, params, convert.ToString(profile, armGeoRedundantBackup))
		cf[tfGeoRedundantBackupEnabled] = strings.EqualFold(strings.ToUpper(status), armStatusEnabled)

		cf[tfBackupRetentionDays] = fn.LookUpFloat64(vars, params, convert.ToString(profile, armBackupRetentionDays))

		cf[tfStorageMB] = fn.LookUpFloat64(vars, params, convert.ToString(profile, armStorageMB))

		status = fn.LookUpString(vars, params, convert.ToString(profile, armSslEnforcement))
		cf[tfSslEnforcementEnabled] = strings.EqualFold(strings.ToUpper(status), armStatusEnabled)
	}
	return cf
}
