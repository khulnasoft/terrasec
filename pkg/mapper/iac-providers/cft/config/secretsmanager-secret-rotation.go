package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/secretsmanager"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SecretRotationRulesBlock holds config for SecretRotationRules
type SecretRotationRulesBlock struct {
	AutomaticallyAfterDays int `json:"automatically_after_days"`
}

// SecretsManagerSecretRotationConfig holds config for SecretsManagerSecretRotation
type SecretsManagerSecretRotationConfig struct {
	Config
	SecretID          string                     `json:"secret_id"`
	RotationLambdaARN string                     `json:"rotation_lambda_arn"`
	RotationRules     []SecretRotationRulesBlock `json:"rotation_rules"`
}

// GetSecretsManagerSecretRotationConfig returns config for SecretsManagerSecretRotation
func GetSecretsManagerSecretRotationConfig(r *secretsmanager.RotationSchedule) []AWSResourceConfig {
	var rotationRules []SecretRotationRulesBlock
	if r.RotationRules != nil {
		rotationRules = make([]SecretRotationRulesBlock, 1)
		rotationRules[0].AutomaticallyAfterDays = functions.GetVal(r.RotationRules.AutomaticallyAfterDays)
	}

	cf := SecretsManagerSecretRotationConfig{
		SecretID:          r.SecretId,
		RotationLambdaARN: functions.GetVal(r.RotationLambdaARN),
		RotationRules:     rotationRules,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
