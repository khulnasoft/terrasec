

package config

import (
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

const (
	// IamUserLoginProfile represents the subresource aws_iam_user_login_profile for attribute LoginProfile
	IamUserLoginProfile = "LoginProfile"
	// IamUserPolicy represents the subresource aws_iam_user_policy for the attribute policy
	IamUserPolicy = "Policy"
)

// IamUserLoginProfileConfig holds config for aws_iam_user_login_profile
type IamUserLoginProfileConfig struct {
	Config
	PasswordResetRequired bool `json:"password_reset_required"`
}

// IamUserPolicyConfig holds config for aws_iam_user_policy
type IamUserPolicyConfig struct {
	Config
	PolicyName     string `json:"name"`
	PolicyDocument string `json:"policy"`
}

// IamUserConfig holds config for aws_iam_user
type IamUserConfig struct {
	Config
	UserName string `json:"name"`
}

// GetIamUserConfig returns config for aws_iam_user, aws_iam_user_policy, aws_iam_user_login_profile
func GetIamUserConfig(i *iam.User) []AWSResourceConfig {

	resourceConfigs := make([]AWSResourceConfig, 0)

	// add aws_iam_user
	resourceConfigs = append(resourceConfigs, AWSResourceConfig{
		Metadata: i.AWSCloudFormationMetadata,
		Resource: IamUserConfig{
			Config: Config{
				Name: functions.GetVal(i.UserName),
				Tags: i.Tags,
			},
			UserName: functions.GetVal(i.UserName),
		},
	})

	iamLoginProfileConfig := IamUserLoginProfileConfig{
		Config: Config{
			Name: functions.GetVal(i.UserName),
		},
	}
	if i.LoginProfile != nil {
		iamLoginProfileConfig.PasswordResetRequired = functions.GetVal(i.LoginProfile.PasswordResetRequired)
	}

	// add aws_iam_user_login_profile
	resourceConfigs = append(resourceConfigs, AWSResourceConfig{
		Type:     IamUserLoginProfile,
		Name:     functions.GetVal(i.UserName),
		Metadata: i.AWSCloudFormationMetadata,
		Resource: iamLoginProfileConfig,
	})

	// add aws_iam_user_policy
	if i.Policies != nil {
		for j, policy := range i.Policies {
			pc := IamUserPolicyConfig{
				Config: Config{
					Name: policy.PolicyName,
				},
				PolicyName: policy.PolicyName,
			}
			policyDocument, err := json.Marshal(policy.PolicyDocument)
			if err == nil {
				pc.PolicyDocument = string(policyDocument)
			}
			resourceConfigs = append(resourceConfigs, AWSResourceConfig{
				Type: IamUserPolicy,
				// Unique name for each policy used for ID
				Name:     fmt.Sprintf("%s%v", policy.PolicyName, j),
				Resource: pc,
				Metadata: i.AWSCloudFormationMetadata,
			})
		}
	}

	return resourceConfigs
}
