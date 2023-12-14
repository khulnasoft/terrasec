

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/elasticache"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// ElastiCacheClusterConfig holds config for aws_elasticache_cluster
type ElastiCacheClusterConfig struct {
	Config
	AZMode        string `json:"az_mode"`
	Engine        string `json:"engine,omitempty"`
	EngineVersion string `json:"engine_version,omitempty"`
}

// GetElastiCacheClusterConfig returns config for aws_elasticache_cluster
func GetElastiCacheClusterConfig(e *elasticache.CacheCluster) []AWSResourceConfig {
	cf := ElastiCacheClusterConfig{
		Config: Config{
			Tags: e.Tags,
			Name: functions.GetVal(e.ClusterName),
		},
		AZMode:        functions.GetVal(e.AZMode),
		Engine:        e.Engine,
		EngineVersion: functions.GetVal(e.EngineVersion),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: e.AWSCloudFormationMetadata,
	}}
}
