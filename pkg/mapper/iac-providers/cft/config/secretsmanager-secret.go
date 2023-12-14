package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/secretsmanager"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SecretsManagerSecretConfig holds config for aws_secretsmanager_secret
type SecretsManagerSecretConfig struct {
	Config
	KmsKeyID string `json:"kms_key_id,omitempty"`
}

// GetSecretsManagerSecretConfig returns config for aws_secretsmanager_secret
func GetSecretsManagerSecretConfig(s *secretsmanager.Secret) []AWSResourceConfig {
	cf := SecretsManagerSecretConfig{
		Config: Config{
			Tags: s.Tags,
			Name: functions.GetVal(s.Name),
		},
		KmsKeyID: functions.GetVal(s.KmsKeyId),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: s.AWSCloudFormationMetadata,
	}}
}
