package config

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armVersion             = "version"
	armStorageProfile      = "storageProfile"
	armStorageMB           = "storageMB"
	armBackupRetentionDays = "backupRetentionDays"
	armGeoRedundantBackup  = "geoRedundantBackup"
	armSslEnforcement      = "sslEnforcement"
	armStatusEnabled       = "ENABLED"
)

const (
	tfSkuName                   = "sku_name"
	tfStorageMB                 = "storage_mb"
	tfVersion                   = "version"
	tfBackupRetentionDays       = "backup_retention_days"
	tfGeoRedundantBackupEnabled = "geo_redundant_backup_enabled"
	tfSslEnforcementEnabled     = "ssl_enforcement_enabled"
)

// MySQLServerConfig returns config for azurerm_mysql_server
func MySQLServerConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation:                   fn.LookUpString(nil, params, r.Location),
		tfName:                       fn.LookUpString(nil, params, r.Name),
		tfSkuName:                    fn.LookUpString(vars, params, r.SKU.Name),
		tfTags:                       r.Tags,
		tfVersion:                    fn.LookUpString(vars, params, convert.ToString(r.Properties, armVersion)),
		tfAdministratorLogin:         fn.LookUpString(vars, params, convert.ToString(r.Properties, armAdministratorLogin)),
		tfAdministratorLoginPassword: fn.LookUpString(vars, params, convert.ToString(r.Properties, armAdministratorLoginPassword)),
	}

	profile := convert.ToMap(r.Properties, armStorageProfile)
	cf[tfStorageMB] = fn.LookUpFloat64(vars, params, convert.ToString(profile, armStorageMB))

	cf[tfBackupRetentionDays] = convert.ToFloat64(profile, armBackupRetentionDays)

	status := strings.ToUpper(convert.ToString(profile, armGeoRedundantBackup))
	cf[tfGeoRedundantBackupEnabled] = strings.EqualFold(status, armStatusEnabled)

	status = strings.ToUpper(convert.ToString(r.Properties, armSslEnforcement))
	cf[tfSslEnforcementEnabled] = strings.EqualFold(status, armStatusEnabled)
	return cf
}
