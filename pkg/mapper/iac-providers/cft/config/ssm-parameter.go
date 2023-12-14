

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ssm"
)

// SSMParameterConfig holds config for SSMParameter
type SSMParameterConfig struct {
	Config
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	Tier           string `json:"tier"`
	Policies       string `json:"policies"`
	AllowedPattern string `json:"allowed_pattern"`
}

// GetSSMParameterConfig returns config for SSM Parameter
func GetSSMParameterConfig(b *ssm.Parameter) []AWSResourceConfig {
	cf := SSMParameterConfig{
		Config: Config{
			Name: *b.Name,
			Tags: b.Tags,
		},
		Name:  *b.Name,
		Type:  b.Type,
		Value: b.Value,
	}
	if b.Description != nil {
		cf.Description = *b.Description
	}
	if b.Tier != nil {
		cf.Tier = *b.Tier
	}
	if b.Policies != nil {
		cf.Policies = *b.Policies
	}
	if b.AllowedPattern != nil {
		cf.AllowedPattern = *b.AllowedPattern
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: b.AWSCloudFormationMetadata,
	}}
}
