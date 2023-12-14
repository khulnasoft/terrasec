

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// NatGatewayConfig holds config for aws_nat_gateway
type NatGatewayConfig struct {
	Config
	AllocationID     string `json:"allocation_id"`
	ConnectivityType string `json:"connectivity_type"`
	SubnetID         string `json:"subnet_id"`
}

// GetNatGatewayConfig returns config for aws_nat_gateway
func GetNatGatewayConfig(e *ec2.NatGateway) []AWSResourceConfig {
	cf := NatGatewayConfig{
		Config: Config{
			Tags: e.Tags,
		},
		AllocationID:     functions.GetVal(e.AllocationId),
		ConnectivityType: functions.GetVal(e.ConnectivityType),
		SubnetID:         e.SubnetId,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: e.AWSCloudFormationMetadata,
	}}
}
