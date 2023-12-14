package config

import (
	"encoding/json"

	"github.com/awslabs/goformation/v7/cloudformation/sqs"
)

// SqsQueuePolicyConfig holds config for SqsQueuePolicy
type SqsQueuePolicyConfig struct {
	Config
	QueueURL string `json:"queue_url"`
	Policy   string `json:"policy"`
}

// GetSqsQueuePolicyConfig returns config for SqsQueuePolicy
func GetSqsQueuePolicyConfig(p *sqs.QueuePolicy) []AWSResourceConfig {
	policyDoc, _ := json.Marshal(p.PolicyDocument)

	cflist := make([]SqsQueuePolicyConfig, len(p.Queues))
	resourcelist := make([]AWSResourceConfig, len(p.Queues))

	for i := range cflist {
		cflist[i].Config.Name = p.Queues[i]
		cflist[i].QueueURL = p.Queues[i]
		cflist[i].Policy = string(policyDoc)

		resourcelist[i].Resource = cflist[i]
		resourcelist[i].Metadata = p.AWSCloudFormationMetadata
	}

	return resourcelist
}
