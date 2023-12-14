

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ecr"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// EcrRepositoryConfig holds config for aws_ecr_repository
type EcrRepositoryConfig struct {
	Config
	ImageScanningConfiguration []ImageScanningConfigurationBlock `json:"image_scanning_configuration"`
	AERP                       interface{}                       `json:"aws_ecr_repository_policy,omitempty"`
}

// ImageScanningConfigurationBlock holds config for image_scanning_configuration attribute
type ImageScanningConfigurationBlock struct {
	ScanOnPush bool `json:"scan_on_push"`
}

// GetEcrRepositoryConfig returns config for aws_ecr_repository
func GetEcrRepositoryConfig(r *ecr.Repository) []AWSResourceConfig {
	var imageScanningConfiguration []ImageScanningConfigurationBlock
	if r.ImageScanningConfiguration != nil {
		imageScanningConfiguration = make([]ImageScanningConfigurationBlock, 1)
		imageScanningConfiguration[0].ScanOnPush = functions.GetVal(r.ImageScanningConfiguration.ScanOnPush)
	}

	cf := EcrRepositoryConfig{
		Config: Config{
			Tags: r.Tags,
			Name: functions.GetVal(r.RepositoryName),
		},
		ImageScanningConfiguration: imageScanningConfiguration,
	}

	cf.AERP = r.RepositoryPolicyText
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
