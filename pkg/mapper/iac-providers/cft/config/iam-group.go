

package config

import (
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

const (
	// IamGroupPolicy represents the sub-resource aws_iam_group_policy for attribute Policy
	IamGroupPolicy = "Policy"
)

// IamGroupPolicyConfig holds config for aws_iam_group_policy
type IamGroupPolicyConfig struct {
	Config
	PolicyName     string `json:"name"`
	PolicyDocument string `json:"policy"`
}

// IamGroupConfig holds config for aws_iam_group
type IamGroupConfig struct {
	Config
	Name string `json:"name"`
}

// GetIamGroupConfig returns config for aws_iam_group_policy
func GetIamGroupConfig(r *iam.Group) []AWSResourceConfig {
	// aws_iam_role_policy as a SubResource
	policyConfigs := make([]AWSResourceConfig, 0)
	if r.Policies != nil {
		for i, policy := range r.Policies {
			pc := IamGroupPolicyConfig{
				Config: Config{
					Name: policy.PolicyName,
				},
				PolicyName: policy.PolicyName,
			}
			policyDocument, err := json.Marshal(policy.PolicyDocument)
			if err == nil {
				pc.PolicyDocument = string(policyDocument)
			}
			policyConfigs = append(policyConfigs, AWSResourceConfig{
				Type: IamGroupPolicy,
				// Unique name for each policy used for ID
				Name:     fmt.Sprintf("%s%v", policy.PolicyName, i),
				Resource: pc,
				Metadata: r.AWSCloudFormationMetadata,
			})
		}
	}

	groupConfig := IamGroupConfig{
		Config: Config{
			Name: functions.GetVal(r.GroupName),
		},
		Name: functions.GetVal(r.GroupName),
	}

	var groupPolicyConfig AWSResourceConfig
	groupPolicyConfig.Resource = groupConfig
	groupPolicyConfig.Metadata = r.AWSCloudFormationMetadata
	policyConfigs = append(policyConfigs, groupPolicyConfig)

	return policyConfigs
}
