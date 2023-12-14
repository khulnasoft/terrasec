

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	tfScope     = "scope"
	tfLockLevel = "lock_level"
	tfNotes     = "notes"
)

// ManagementLockConfig returns config for azurerm_management_lock
func ManagementLockConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:  fn.LookUpString(nil, params, r.Location),
		tfName:      fn.LookUpString(nil, params, r.Name),
		tfTags:      r.Tags,
		tfScope:     fn.LookUpString(vars, params, r.Scope),
		tfLockLevel: convert.ToString(r.Properties, armLevel),
		tfNotes:     convert.ToString(r.Properties, tfNotes),
	}
}
