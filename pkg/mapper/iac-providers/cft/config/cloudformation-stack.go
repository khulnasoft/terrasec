

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/cloudformation"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// CloudFormationStackConfig holds config for aws_cloudformation_stack
type CloudFormationStackConfig struct {
	Config
	TemplateURL      string            `json:"template_url"`
	NotificationARNs interface{}       `json:"notification_arns"`
	Parameters       map[string]string `json:"-"`
	TemplateData     []byte            `json:"-"`
}

// GetCloudFormationStackConfig returns config for aws_cloudformation_stack
func GetCloudFormationStackConfig(s *cloudformation.Stack) []AWSResourceConfig {
	cf := CloudFormationStackConfig{
		Config:           Config{Tags: s.Tags},
		TemplateURL:      "",
		NotificationARNs: nil,
		TemplateData:     []byte{},
	}

	if s.NotificationARNs != nil {
		cf.NotificationARNs = s.NotificationARNs
	}

	// Add and resolve template URL
	if len(s.TemplateURL) > 0 {
		cf.TemplateURL = s.TemplateURL

		templateData, err := fn.DownloadBucketObj(s.TemplateURL)
		if err == nil {
			cf.TemplateData = templateData
		}
	}

	// Add Parameters for propagation to the nested Stack
	if s.Parameters != nil {
		cf.Parameters = s.Parameters
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: s.AWSCloudFormationMetadata,
	}}
}
