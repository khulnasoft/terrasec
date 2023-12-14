

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/sagemaker"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SagemakerNotebookInstanceConfig holds config for SagemakerNotebookInstance
type SagemakerNotebookInstanceConfig struct {
	Config
	Name                 string `json:"name"`
	RoleARN              string `json:"role_arn"`
	InstanceType         string `json:"instance_type"`
	KMSKeyID             string `json:"kms_key_id"`
	DirectInternetAccess string `json:"direct_internet_access"`
	RootAccess           string `json:"root_access"`
}

// GetSagemakerNotebookInstanceConfig returns config for SagemakerNotebookInstance
func GetSagemakerNotebookInstanceConfig(n *sagemaker.NotebookInstance) []AWSResourceConfig {
	cf := SagemakerNotebookInstanceConfig{
		Config: Config{
			Name: functions.GetVal(n.NotebookInstanceName),
			Tags: n.Tags,
		},
		Name:                 functions.GetVal(n.NotebookInstanceName),
		RoleARN:              n.RoleArn,
		InstanceType:         n.InstanceType,
		KMSKeyID:             functions.GetVal(n.KmsKeyId),
		DirectInternetAccess: functions.GetVal(n.DirectInternetAccess),
		RootAccess:           functions.GetVal(n.RootAccess),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: n.AWSCloudFormationMetadata,
	}}
}
