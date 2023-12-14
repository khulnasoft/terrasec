package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/sagemaker"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// ImageConfigBlock holds config for ImageConfig
type ImageConfigBlock struct {
	RepositoryAccessMode string `json:"repository_access_mode"`
}

// ContainerBlock holds config for Container
type ContainerBlock struct {
	Image             string             `json:"image"`
	Mode              string             `json:"mode"`
	ModelDataURL      string             `json:"model_data_url"`
	ContainerHostname string             `json:"container_hostname"`
	Environment       interface{}        `json:"environment"`
	ImageConfig       []ImageConfigBlock `json:"image_config"`
}

// SagemakerModelConfig holds config for SagemakerModel
type SagemakerModelConfig struct {
	Config
	Name             string           `json:"name"`
	ExecutionRoleARN string           `json:"execution_role_arn"`
	Container        []ContainerBlock `json:"container"`
	PrimaryContainer []ContainerBlock `json:"primary_container"`
}

// GetSagemakerModelConfig returns config for SagemakerModel
func GetSagemakerModelConfig(m *sagemaker.Model) []AWSResourceConfig {
	var containerBlock []ContainerBlock
	if m.Containers != nil {
		containerBlock = make([]ContainerBlock, len(m.Containers))
		for i, container := range m.Containers {
			containerBlock[i] = getContainer(container)
		}
	}

	var primaryContainer []ContainerBlock
	if m.PrimaryContainer != nil {
		primaryContainer = make([]ContainerBlock, 1)
		primaryContainer[0] = getContainer(*m.PrimaryContainer)
	}

	cf := SagemakerModelConfig{
		Config: Config{
			Name: functions.GetVal(m.ModelName),
			Tags: m.Tags,
		},
		Name:             functions.GetVal(m.ModelName),
		ExecutionRoleARN: m.ExecutionRoleArn,
		Container:        containerBlock,
		PrimaryContainer: primaryContainer,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: m.AWSCloudFormationMetadata,
	}}
}

func getContainer(gftContainer sagemaker.Model_ContainerDefinition) ContainerBlock {
	var container ContainerBlock

	container.Image = functions.GetVal(gftContainer.Image)
	container.Mode = functions.GetVal(gftContainer.Mode)
	container.ModelDataURL = functions.GetVal(gftContainer.ModelDataUrl)
	container.ContainerHostname = functions.GetVal(gftContainer.ContainerHostname)
	container.Environment = gftContainer.Environment

	if gftContainer.ImageConfig != nil {
		container.ImageConfig = make([]ImageConfigBlock, 1)
		container.ImageConfig[0].RepositoryAccessMode = gftContainer.ImageConfig.RepositoryAccessMode
	}

	return container
}
