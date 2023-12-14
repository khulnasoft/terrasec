package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ecs"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// EcsServiceConfig holds config for aws_ecs_service
type EcsServiceConfig struct {
	Config
	IamRole string `json:"iam_role"`
}

// GetEcsServiceConfig returns config for aws_ecs_service
func GetEcsServiceConfig(c *ecs.Service) []AWSResourceConfig {
	cf := EcsServiceConfig{
		Config: Config{
			Name: functions.GetVal(c.ServiceName),
			Tags: c.Tags,
		},
		IamRole: functions.GetVal(c.Role),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
