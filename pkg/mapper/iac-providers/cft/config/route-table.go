package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
)

// RouteTableConfig holds config for aws_route_table
type RouteTableConfig struct {
	Config
	VpcID string `json:"vpc_id"`
}

// GetRouteTableConfig returns config for aws_route_table
func GetRouteTableConfig(e *ec2.RouteTable) []AWSResourceConfig {
	cf := RouteTableConfig{
		Config: Config{
			Tags: e.Tags,
		},
		VpcID: e.VpcId,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: e.AWSCloudFormationMetadata,
	}}
}
