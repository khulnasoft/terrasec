

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/sqs"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SqsQueueConfig holds config for SqsQueue
type SqsQueueConfig struct {
	Config
	Name                         string `json:"name"`
	KmsMasterKeyID               string `json:"kms_master_key_id"`
	KmsDataKeyReusePeriodSeconds int    `json:"kms_data_key_reuse_period_seconds"`
	MessageRetentionSeconds      int    `json:"message_retention_seconds"`
}

// GetSqsQueueConfig returns config for SqsQueue
func GetSqsQueueConfig(q *sqs.Queue) []AWSResourceConfig {
	cf := SqsQueueConfig{
		Config: Config{
			Name: functions.GetVal(q.QueueName),
		},
		Name:                         functions.GetVal(q.QueueName),
		KmsMasterKeyID:               functions.GetVal(q.KmsMasterKeyId),
		KmsDataKeyReusePeriodSeconds: functions.GetVal(q.KmsDataKeyReusePeriodSeconds),
		MessageRetentionSeconds:      functions.GetVal(q.MessageRetentionPeriod),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: q.AWSCloudFormationMetadata,
	}}
}
