

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/logs"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// LogCloudWatchGroupConfig holds config for aws_cloudwatch_log_group
type LogCloudWatchGroupConfig struct {
	Config
	LogGroupName    string `json:"name"`
	KmsKeyID        string `json:"kms_key_id,omitempty"`
	RetentionInDays int    `json:"retention_in_days"`
}

// GetLogCloudWatchGroupConfig returns config for aws_cloudwatch_log_group
func GetLogCloudWatchGroupConfig(r *logs.LogGroup) []AWSResourceConfig {
	cf := LogCloudWatchGroupConfig{
		Config: Config{
			Name: functions.GetVal(r.LogGroupName),
		},
		LogGroupName:    functions.GetVal(r.LogGroupName),
		KmsKeyID:        functions.GetVal(r.KmsKeyId),
		RetentionInDays: functions.GetVal(r.RetentionInDays),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
