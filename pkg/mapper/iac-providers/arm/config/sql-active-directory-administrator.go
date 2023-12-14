

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armLogin    = "login"
	armSid      = "sid"
	armTenantID = "tenantId"
)
const (
	tfLogin    = "login"
	tfObjectID = "object_id"
)

// SQLActiveDirectoryAdministratorConfig returns config for azurerm_sql_active_directory_administrator
func SQLActiveDirectoryAdministratorConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfLogin:    fn.LookUpString(vars, params, convert.ToString(r.Properties, armLogin)),
		tfObjectID: fn.LookUpString(vars, params, convert.ToString(r.Properties, armSid)),
		tfTenantID: fn.LookUpString(vars, params, convert.ToString(r.Properties, armTenantID)),
	}
}
