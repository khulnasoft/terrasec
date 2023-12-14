

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/rds"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// RDSClusterConfig holds config for aws_rds_cluster
type RDSClusterConfig struct {
	Config
	BackupRetentionPeriod int  `json:"backup_retention_period,omitempty"`
	StorageEncrypted      bool `json:"storage_encrypted"`
}

// GetRDSClusterConfig returns config for aws_rds_cluster
func GetRDSClusterConfig(c *rds.DBCluster) []AWSResourceConfig {
	cf := RDSClusterConfig{
		Config: Config{
			Name: functions.GetVal(c.DatabaseName),
			Tags: c.Tags,
		},
		BackupRetentionPeriod: functions.GetVal(c.BackupRetentionPeriod),
		StorageEncrypted:      functions.GetVal(c.StorageEncrypted),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
