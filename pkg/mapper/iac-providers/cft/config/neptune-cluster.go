

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/neptune"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// NeptuneClusterConfig holds config for aws_neptune_cluster
type NeptuneClusterConfig struct {
	Config
	EnableCloudwatchLogsExports []string `json:"enable_cloudwatch_logs_exports,omitempty"`
	StorageEncrypted            bool     `json:"storage_encrypted,omitempty"`
}

// GetNeptuneClusterConfig returns config for aws_neptune_cluster
func GetNeptuneClusterConfig(d *neptune.DBCluster) []AWSResourceConfig {
	cf := NeptuneClusterConfig{
		Config: Config{
			Tags: d.Tags,
		},
		StorageEncrypted:            functions.GetVal(d.StorageEncrypted),
		EnableCloudwatchLogsExports: d.EnableCloudwatchLogsExports,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: d.AWSCloudFormationMetadata,
	}}
}
