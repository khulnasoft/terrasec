

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/route53"
)

// Route53RecordConfig holds config for aws_route53_record
type Route53RecordConfig struct {
	Config
	ResourceRecords []string `json:"records"`
}

// GetRoute53RecordConfig returns config for aws_route53_record
func GetRoute53RecordConfig(r *route53.RecordSet) []AWSResourceConfig {
	cf := Route53RecordConfig{
		Config: Config{
			Name: r.Name,
		},
		ResourceRecords: r.ResourceRecords,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
