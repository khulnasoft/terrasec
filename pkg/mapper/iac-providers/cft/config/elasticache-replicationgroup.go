

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/elasticache"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// ElastiCacheReplicationGroupConfig holds config for aws_elasticache_replication_group
type ElastiCacheReplicationGroupConfig struct {
	Config
	AtRestEncryptionEnabled  bool `json:"at_rest_encryption_enabled,omitempty"`
	TransitEncryptionEnabled bool `json:"transit_encryption_enabled,omitempty"`
}

// GetElastiCacheReplicationGroupConfig returns config for aws_elasticache_replication_group
func GetElastiCacheReplicationGroupConfig(r *elasticache.ReplicationGroup) []AWSResourceConfig {
	cf := ElastiCacheReplicationGroupConfig{
		Config: Config{
			Tags: r.Tags,
		},
		AtRestEncryptionEnabled:  functions.GetVal(r.AtRestEncryptionEnabled),
		TransitEncryptionEnabled: functions.GetVal(r.TransitEncryptionEnabled),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
