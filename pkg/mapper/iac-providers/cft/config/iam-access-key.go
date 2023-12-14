

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// IamAccessKeyConfig holds config for aws_iam_access_key
type IamAccessKeyConfig struct {
	Config
	UserName string `json:"user"`
	Status   string `json:"status"`
}

// GetIamAccessKeyConfig returns config for aws_iam_access_key
func GetIamAccessKeyConfig(r *iam.AccessKey) []AWSResourceConfig {
	cf := IamAccessKeyConfig{
		Config: Config{
			Name: r.UserName,
		},
		UserName: r.UserName,
		Status:   functions.GetVal(r.Status),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
