

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// EbsVolumeConfig holds config for aws_ebs_volume
type EbsVolumeConfig struct {
	Config
	Encrypted bool   `json:"encrypted,omitempty"`
	KmsKeyID  string `json:"kms_key_id,omitempty"`
}

// GetEbsVolumeConfig returns config for aws_ebs_volume
func GetEbsVolumeConfig(v *ec2.Volume) []AWSResourceConfig {
	cf := EbsVolumeConfig{
		Config: Config{
			Tags: v.Tags,
		},
		Encrypted: functions.GetVal(v.Encrypted),
		KmsKeyID:  functions.GetVal(v.KmsKeyId),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: v.AWSCloudFormationMetadata,
	}}
}
