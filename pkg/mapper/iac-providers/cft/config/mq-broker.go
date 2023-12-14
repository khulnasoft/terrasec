package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/amazonmq"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// MqBrokerConfig holds config for aws_mq_broker
type MqBrokerConfig struct {
	Logs interface{} `json:"logs,omitempty"`
	Config
	PubliclyAccessible bool `json:"publicly_accessible"`
}

// GetMqBorkerConfig returns config for aws_mq_broker
func GetMqBorkerConfig(c *amazonmq.Broker) []AWSResourceConfig {
	cf := MqBrokerConfig{
		Config: Config{
			Name: c.BrokerName,
			Tags: c.Tags,
		},
		PubliclyAccessible: c.PubliclyAccessible,
	}
	if c.Logs != nil {
		log := make(map[string]bool)
		if functions.GetVal(c.Logs.Audit) {
			log["audit"] = true
		} else {
			log["audit"] = false
		}
		if functions.GetVal(c.Logs.General) {
			log["general"] = true
		} else {
			log["general"] = false
		}
		cf.Logs = []map[string]bool{log}
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
