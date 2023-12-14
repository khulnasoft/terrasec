

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/applicationautoscaling"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// AppAutoScalingPolicyConfig holds config for AppAutoScalingPolicy
type AppAutoScalingPolicyConfig struct {
	Config
	Name              string `json:"name"`
	PolicyType        string `json:"policy_type"`
	ResourceID        string `json:"resource_id"`
	ScalableDimension string `json:"scalable_dimension"`
	ServiceNamespace  string `json:"service_namespace"`
}

// GetAppAutoScalingPolicyConfig returns config for AppAutoScalingPolicy
func GetAppAutoScalingPolicyConfig(a *applicationautoscaling.ScalingPolicy) []AWSResourceConfig {
	cf := AppAutoScalingPolicyConfig{
		Config: Config{
			Name: a.PolicyName,
		},
		Name:              a.PolicyName,
		PolicyType:        a.PolicyType,
		ResourceID:        functions.GetVal(a.ResourceId),
		ScalableDimension: functions.GetVal(a.ScalableDimension),
		ServiceNamespace:  functions.GetVal(a.ServiceNamespace),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: a.AWSCloudFormationMetadata,
	}}
}
