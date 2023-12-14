

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armEncryptionSettingsCollection = "encryptionSettingsCollection"
	armCreationData                 = "creationData"
	armCreateOption                 = "createOption"
	armDiskSizeGB                   = "diskSizeGB"
	armSourceResourceID             = "sourceResourceId"
)

const (
	tfCreateOption       = "create_option"
	tfDiskSizeGB         = "disk_size_gb"
	tfSourceResourceID   = "source_resource_id"
	tfStorageAccountType = "storage_account_type"
	tfEncryptionSettings = "encryption_settings"
)

// ManagedDiskConfig returns config for azurerm_managed_disk.
func ManagedDiskConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation:           fn.LookUpString(nil, params, r.Location),
		tfName:               fn.LookUpString(nil, params, r.Name),
		tfTags:               r.Tags,
		tfStorageAccountType: r.SKU.Name,
		tfEncryptionSettings: convert.ToMap(r.Properties, armEncryptionSettingsCollection),
		tfDiskSizeGB:         fn.LookUpFloat64(vars, params, convert.ToString(r.Properties, armDiskSizeGB)),
	}

	data := convert.ToMap(r.Properties, armCreationData)
	cf[tfCreateOption] = convert.ToString(data, armCreateOption)
	cf[tfSourceResourceID] = convert.ToString(data, armSourceResourceID)
	return cf
}
