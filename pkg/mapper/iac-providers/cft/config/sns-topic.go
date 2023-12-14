package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/sns"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SnsTopicConfig holds config for SnsTopic
type SnsTopicConfig struct {
	Config
	Name        string `json:"name"`
	KmsMasterID string `json:"kms_master_id"`
}

// GetSnsTopicConfig returns config for SnsTopic
func GetSnsTopicConfig(t *sns.Topic) []AWSResourceConfig {
	cf := SnsTopicConfig{
		Config: Config{
			Name: functions.GetVal(t.TopicName),
		},
		Name:        functions.GetVal(t.TopicName),
		KmsMasterID: functions.GetVal(t.KmsMasterKeyId),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: t.AWSCloudFormationMetadata,
	}}
}
