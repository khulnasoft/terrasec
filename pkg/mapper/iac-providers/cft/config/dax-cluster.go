package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/dax"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// DaxClusterConfig holds config for aws_dax_cluster
type DaxClusterConfig struct {
	Config
	ServerSideEncryption []SSE `json:"server_side_encryption"`
}

// GetDaxClusterConfig returns config for aws_dax_cluster
func GetDaxClusterConfig(t *dax.Cluster) []AWSResourceConfig {
	cf := DaxClusterConfig{
		Config: Config{
			Tags: t.Tags,
			Name: functions.GetVal(t.ClusterName),
		},
	}

	if t.SSESpecification != nil {
		cf.ServerSideEncryption = make([]SSE, 1)

		cf.ServerSideEncryption[0].Enabled = functions.GetVal(t.SSESpecification.SSEEnabled)
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: t.AWSCloudFormationMetadata,
	}}
}
