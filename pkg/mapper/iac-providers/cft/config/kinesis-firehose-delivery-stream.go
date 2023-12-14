package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/kinesisfirehose"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// KinesisFirehoseDeliveryStreamConfig holds config for aws_kinesis_firehose_delivery_stream
type KinesisFirehoseDeliveryStreamConfig struct {
	ServerSideEncryption interface{} `json:"server_side_encryption"`
	Config
}

// KinesisFirehoseDeliveryStreamSseConfig holds config for server_side_encryption attribute of aws_kinesis_firehose_delivery_stream
type KinesisFirehoseDeliveryStreamSseConfig struct {
	KeyType string `json:"key_type,omitempty"`
	KeyARN  string `json:"key_arn,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// GetKinesisFirehoseDeliveryStreamConfig returns aws_kinesis_firehose_delivery_stream
func GetKinesisFirehoseDeliveryStreamConfig(k *kinesisfirehose.DeliveryStream) []AWSResourceConfig {
	cf := KinesisFirehoseDeliveryStreamConfig{
		Config: Config{
			Name: functions.GetVal(k.DeliveryStreamName),
			Tags: k.Tags,
		},
	}
	sseConfig := KinesisFirehoseDeliveryStreamSseConfig{}
	if k.DeliveryStreamEncryptionConfigurationInput != nil {
		sseConfig.Enabled = true
		sseConfig.KeyType = k.DeliveryStreamEncryptionConfigurationInput.KeyType
		sseConfig.KeyARN = functions.GetVal(k.DeliveryStreamEncryptionConfigurationInput.KeyARN)
	}
	cf.ServerSideEncryption = []KinesisFirehoseDeliveryStreamSseConfig{sseConfig}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: k.AWSCloudFormationMetadata,
	}}
}
