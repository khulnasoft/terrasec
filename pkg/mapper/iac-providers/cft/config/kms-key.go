package config

import (
	"encoding/json"

	"github.com/awslabs/goformation/v7/cloudformation/kms"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// KmsKeyConfig holds config for aws_kms_key
type KmsKeyConfig struct {
	Config
	Description         string `json:"description"`
	KeyPolicy           string `json:"policy"`
	PendingWindowInDays int    `json:"deletion_window_in_days"`
	Enabled             bool   `json:"is_enabled"`
	EnableKeyRotation   bool   `json:"enable_key_rotation"`
}

// GetKmsKeyConfig returns config for aws_kms_key
func GetKmsKeyConfig(k *kms.Key) []AWSResourceConfig {
	cf := KmsKeyConfig{
		Config: Config{
			Tags: k.Tags,
		},
		Enabled:             functions.GetVal(k.Enabled),
		EnableKeyRotation:   functions.GetVal(k.EnableKeyRotation),
		PendingWindowInDays: functions.GetVal(k.PendingWindowInDays),
	}

	keyPolicy, err := json.Marshal(k.KeyPolicy)
	if err == nil {
		cf.KeyPolicy = string(keyPolicy)
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: k.AWSCloudFormationMetadata,
	}}
}
