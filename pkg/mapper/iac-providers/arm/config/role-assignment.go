package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armLevel            = "level"
	armPrincipalID      = "principalId"
	armRoleDefinitionID = "roleDefinitionId"
)

const (
	tfPrincipalID      = "principal_id"
	tfRoleDefinitionID = "role_definition_id"
)

// RoleAssignmentConfig returns config for azurerm_role_assignment
func RoleAssignmentConfig(r types.Resource, vars, params map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		tfLocation:         fn.LookUpString(nil, params, r.Location),
		tfName:             fn.LookUpString(nil, params, r.Name),
		tfTags:             r.Tags,
		tfScope:            fn.LookUpString(vars, params, r.Scope),
		tfLockLevel:        convert.ToString(r.Properties, armLevel),
		tfPrincipalID:      convert.ToString(r.Properties, armPrincipalID),
		tfRoleDefinitionID: fn.LookUpString(vars, params, convert.ToString(r.Properties, armRoleDefinitionID)),
	}
}
