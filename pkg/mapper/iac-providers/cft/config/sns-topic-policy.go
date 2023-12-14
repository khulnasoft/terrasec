package config

import (
	"encoding/json"

	"github.com/awslabs/goformation/v7/cloudformation/sns"
)

// SnsTopicPolicyConfig holds config for SnsTopicPolicy
type SnsTopicPolicyConfig struct {
	Config
	ARN    string `json:"arn"`
	Policy string `json:"policy"`
}

// GetSnsTopicPolicyConfig returns config for SnsTopicPolicy
func GetSnsTopicPolicyConfig(p *sns.TopicPolicy) []AWSResourceConfig {
	policyDoc, _ := json.Marshal(p.PolicyDocument)

	cflist := make([]SnsTopicPolicyConfig, len(p.Topics))
	resourcelist := make([]AWSResourceConfig, len(p.Topics))

	for i := range p.Topics {
		cflist[i].Config.Name = p.Topics[i]
		cflist[i].ARN = p.Topics[i]
		cflist[i].Policy = string(policyDoc)

		resourcelist[i].Resource = cflist[i]
		resourcelist[i].Metadata = p.AWSCloudFormationMetadata
	}

	return resourcelist
}
