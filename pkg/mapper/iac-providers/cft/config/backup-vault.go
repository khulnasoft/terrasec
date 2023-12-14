package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/backup"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// BackupVaultConfig holds config for BackupVault
type BackupVaultConfig struct {
	Config
	Name      string `json:"name"`
	KMSKeyARN string `json:"kms_key_arn"`
}

// GetBackupVaultConfig returns config for BackupVault
func GetBackupVaultConfig(b *backup.BackupVault) []AWSResourceConfig {
	cf := BackupVaultConfig{
		Config: Config{
			Name: b.BackupVaultName,
			Tags: b.BackupVaultTags,
		},
		Name:      b.BackupVaultName,
		KMSKeyARN: functions.GetVal(b.EncryptionKeyArn),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: b.AWSCloudFormationMetadata,
	}}
}
