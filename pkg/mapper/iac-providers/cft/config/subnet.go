package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// SubnetConfig holds config for aws_subnet
type SubnetConfig struct {
	Config
	AssignIpv6AddressOnCreation bool   `json:"assign_ipv6_address_on_creation"`
	AvailabilityZone            string `json:"availability_zone"`
	CidrBlock                   string `json:"cidr_block"`
	Ipv6CidrBlock               string `json:"ipv6_cidr_block"`
	MapPublicIPOnLaunch         bool   `json:"map_public_ip_on_launch"`
	OutpostArn                  string `json:"outpost_arn"`
	VpcID                       string `json:"vpc_id"`
}

// GetSubnetConfig returns config for aws_subnet
func GetSubnetConfig(e *ec2.Subnet) []AWSResourceConfig {
	cf := SubnetConfig{
		Config: Config{
			Tags: e.Tags,
		},
		AssignIpv6AddressOnCreation: functions.GetVal(e.AssignIpv6AddressOnCreation),
		AvailabilityZone:            functions.GetVal(e.AvailabilityZone),
		CidrBlock:                   functions.GetVal(e.CidrBlock),
		Ipv6CidrBlock:               functions.GetVal(e.Ipv6CidrBlock),
		MapPublicIPOnLaunch:         functions.GetVal(e.MapPublicIpOnLaunch),
		OutpostArn:                  functions.GetVal(e.OutpostArn),
		VpcID:                       e.VpcId,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: e.AWSCloudFormationMetadata,
	}}
}
