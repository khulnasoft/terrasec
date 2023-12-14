package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/docdb"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// DocDBClusterConfig holds config for aws_docdb_cluster
type DocDBClusterConfig struct {
	Config
	KmsKeyID                    string   `json:"kms_key_id,omitempty"`
	EnableCloudwatchLogsExports []string `json:"enabled_cloudwatch_logs_exports"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
}

// GetDocDBConfig returns config for aws_docdb_cluster
func GetDocDBConfig(d *docdb.DBCluster) []AWSResourceConfig {
	cf := DocDBClusterConfig{
		Config: Config{
			Tags: d.Tags,
		},
		KmsKeyID:                    functions.GetVal(d.KmsKeyId),
		StorageEncrypted:            functions.GetVal(d.StorageEncrypted),
		EnableCloudwatchLogsExports: d.EnableCloudwatchLogsExports,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: d.AWSCloudFormationMetadata,
	}}
}
