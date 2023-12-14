package config

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const publicAccess = "publicAccess"
const tfContainerAccessType = "container_access_type"

// StorageContainerConfig returns config for azurerm_storage_container
func StorageContainerConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	access := fn.LookUpString(nil, params, convert.ToString(r.Properties, publicAccess))
	if strings.ToUpper(access) == "NONE" {
		access = "private"
	}
	cf[tfContainerAccessType] = access
	return cf
}
