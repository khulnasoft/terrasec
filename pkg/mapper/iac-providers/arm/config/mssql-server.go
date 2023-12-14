

package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armAdministratorLogin         = "administratorLogin"
	armAdministratorLoginPassword = "administratorLoginPassword"
	armMinimumTLSVersion          = "minimalTlsVersion"
)

const (
	tfAdministratorLogin         = "administrator_login"
	tfAdministratorLoginPassword = "administrator_login_password"
	tfMinimumTLSVersion          = "minimum_tls_version"
)

// MSSQLServerConfig returns config for azurerm_mssql_server
func MSSQLServerConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:                   fn.LookUpString(nil, params, r.Location),
		tfName:                       fn.LookUpString(nil, params, r.Name),
		tfTags:                       r.Tags,
		tfAdministratorLogin:         fn.LookUpString(vars, params, convert.ToString(r.Properties, armAdministratorLogin)),
		tfAdministratorLoginPassword: fn.LookUpString(vars, params, convert.ToString(r.Properties, armAdministratorLoginPassword)),
		tfMinimumTLSVersion:          fn.LookUpString(vars, params, convert.ToString(r.Properties, armMinimumTLSVersion)),
	}
}
