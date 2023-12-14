package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/cognito"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// PasswordPolicyBlock holds config for PasswordPolicy
type PasswordPolicyBlock struct {
	MinimumLength                 int  `json:"minimum_length"`
	RequireLowercase              bool `json:"require_lowercase"`
	RequireUppercase              bool `json:"require_uppercase"`
	RequireNumbers                bool `json:"require_numbers"`
	RequireSymbols                bool `json:"require_symbols"`
	TemporaryPasswordValidityDays int  `json:"temporary_password_validity_days"`
}

// CognitoUserPoolConfig holds config for CognitoUserPool
type CognitoUserPoolConfig struct {
	Config
	Name           string                `json:"name"`
	PasswordPolicy []PasswordPolicyBlock `json:"password_policy"`
}

// GetCognitoUserPoolConfig returns config for CognitoUserPool
func GetCognitoUserPoolConfig(u *cognito.UserPool) []AWSResourceConfig {
	var passwordPolicy []PasswordPolicyBlock
	if u.Policies != nil && u.Policies.PasswordPolicy != nil {
		passwordPolicy = make([]PasswordPolicyBlock, 1)

		passwordPolicy[0].MinimumLength = functions.GetVal(u.Policies.PasswordPolicy.MinimumLength)
		passwordPolicy[0].RequireLowercase = functions.GetVal(u.Policies.PasswordPolicy.RequireLowercase)
		passwordPolicy[0].RequireUppercase = functions.GetVal(u.Policies.PasswordPolicy.RequireUppercase)
		passwordPolicy[0].RequireNumbers = functions.GetVal(u.Policies.PasswordPolicy.RequireNumbers)
		passwordPolicy[0].RequireSymbols = functions.GetVal(u.Policies.PasswordPolicy.RequireSymbols)
		passwordPolicy[0].TemporaryPasswordValidityDays = functions.GetVal(u.Policies.PasswordPolicy.TemporaryPasswordValidityDays)
	}

	cf := CognitoUserPoolConfig{
		Config: Config{
			Name: functions.GetVal(u.UserPoolName),
		},
		Name:           functions.GetVal(u.UserPoolName),
		PasswordPolicy: passwordPolicy,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: u.AWSCloudFormationMetadata,
	}}
}
