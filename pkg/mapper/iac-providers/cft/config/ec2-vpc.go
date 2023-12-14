

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// Ec2VpcConfig holds config for Ec2Vpc
type Ec2VpcConfig struct {
	Config
	CIDRBlock          string `json:"cidr_block"`
	EnableDNSSupport   bool   `json:"enable_dns_support"`
	EnableDNSHostnames bool   `json:"enable_dns_hostnames"`
	InstanceTenancy    string `json:"instance_tenancy"`
}

// GetEc2VpcConfig returns config for Ec2Vpc
func GetEc2VpcConfig(v *ec2.VPC) []AWSResourceConfig {
	cf := Ec2VpcConfig{
		Config: Config{
			Tags: v.Tags,
		},
		CIDRBlock:          functions.GetVal(v.CidrBlock),
		EnableDNSSupport:   functions.GetVal(v.EnableDnsSupport),
		EnableDNSHostnames: functions.GetVal(v.EnableDnsHostnames),
		InstanceTenancy:    functions.GetVal(v.InstanceTenancy),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: v.AWSCloudFormationMetadata,
	}}
}
