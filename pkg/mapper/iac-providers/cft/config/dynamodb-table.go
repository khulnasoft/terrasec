package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/dynamodb"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// PITR holds config for point_in_time_recovery block
type PITR struct {
	Enabled bool `json:"enabled"`
}

// SSE holds config for server_side_encryption block
type SSE struct {
	Enabled bool `json:"enabled"`
}

// DynamoDBTableConfig holds config for aws_dynamodb_table
type DynamoDBTableConfig struct {
	Config
	ServerSideEncryption []SSE  `json:"server_side_encryption"`
	PointInTimeRecovery  []PITR `json:"point_in_time_recovery"`
}

// GetDynamoDBTableConfig returns config for aws_dynamodb_table
func GetDynamoDBTableConfig(t *dynamodb.Table) []AWSResourceConfig {
	cf := DynamoDBTableConfig{
		Config: Config{
			Tags: t.Tags,
			Name: functions.GetVal(t.TableName),
		},
	}

	if t.SSESpecification != nil {
		cf.ServerSideEncryption = make([]SSE, 1)

		cf.ServerSideEncryption[0].Enabled = t.SSESpecification.SSEEnabled
	}

	if t.PointInTimeRecoverySpecification != nil {
		cf.PointInTimeRecovery = make([]PITR, 1)

		cf.PointInTimeRecovery[0].Enabled = functions.GetVal(t.PointInTimeRecoverySpecification.PointInTimeRecoveryEnabled)
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: t.AWSCloudFormationMetadata,
	}}
}
