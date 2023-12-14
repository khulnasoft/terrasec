package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
)

// RouteTableAssociationConfig holds config for aws_route_table_association
type RouteTableAssociationConfig struct {
	Config
	RouteTableID string `json:"route_table_id"`
	SubnetID     string `json:"subnet_id"`
}

// GetRouteTableAssociationConfig returns config for aws_route_table_association
func GetRouteTableAssociationConfig(e *ec2.SubnetRouteTableAssociation) []AWSResourceConfig {
	cf := RouteTableAssociationConfig{
		RouteTableID: e.RouteTableId,
		SubnetID:     e.SubnetId,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: e.AWSCloudFormationMetadata,
	}}
}
