package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/cloudtrail"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// CloudTrailConfig holds config for aws_cloudtrail
type CloudTrailConfig struct {
	Config
	IsMultiRegionTrail      interface{} `json:"is_multi_region_trail"`
	KmsKeyID                interface{} `json:"kms_key_id"`
	SnsTopicName            interface{} `json:"sns_topic_name"`
	EnableLogFileValidation interface{} `json:"enable_log_file_validation"`
}

// GetCloudTrailConfig returns config for aws_cloudtrail
func GetCloudTrailConfig(t *cloudtrail.Trail) []AWSResourceConfig {
	cf := CloudTrailConfig{
		Config:                  Config{Tags: t.Tags, Name: functions.GetVal(t.TrailName)},
		EnableLogFileValidation: t.EnableLogFileValidation,
		IsMultiRegionTrail:      t.IsMultiRegionTrail,
	}
	cf.KmsKeyID = functions.GetVal(t.KMSKeyId)
	cf.SnsTopicName = functions.GetVal(t.SnsTopicName)
	return []AWSResourceConfig{{Resource: cf, Metadata: t.AWSCloudFormationMetadata}}
}
