

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armStorageEndpoint            = "storageEndpoint"
	armStorageAccountAccessKey    = "storageAccountAccessKey"
	armIsStorageSecondaryKeyInUse = "isStorageSecondaryKeyInUse"
	armRetentionDays              = "retentionDays"
)

const (
	tfStorageEndpoint                    = "storage_endpoint,omitempty"
	tfStorageAccountAccessKey            = "storage_account_access_key,omitempty"
	tfStorageAccountAccessKeyIsSecondary = "storage_account_access_key_is_secondary,omitempty"
	tfRetentionInDays                    = "retention_in_days,omitempty"
)

// AuditingPolicyConfig returns config for azurerm_mssql_database_extended_auditing_policy
func AuditingPolicyConfig(r types.Resource) map[string]interface{} {
	return map[string]interface{}{
		tfStorageEndpoint:                    convert.ToString(r.Properties, armStorageEndpoint),
		tfStorageAccountAccessKey:            convert.ToString(r.Properties, armStorageAccountAccessKey),
		tfStorageAccountAccessKeyIsSecondary: convert.ToBool(r.Properties, armIsStorageSecondaryKeyInUse),
		tfRetentionInDays:                    convert.ToFloat64(r.Properties, armRetentionDays),
	}
}
