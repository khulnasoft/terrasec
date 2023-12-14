package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/globalaccelerator"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// GlobalAcceleratorConfig holds config for aws_globalaccelerator_accelerator resource
type GlobalAcceleratorConfig struct {
	Config
	Name          string `json:"name"`
	IPAddressType string `json:"ip_address_type"`
	Enabled       bool   `json:"enabled"`
}

// GetGlobalAcceleratorConfig returns config for aws_globalaccelerator_accelerator resource
func GetGlobalAcceleratorConfig(g *globalaccelerator.Accelerator) []AWSResourceConfig {
	cf := GlobalAcceleratorConfig{
		Config: Config{
			Name: g.Name,
			Tags: g.Tags,
		},
		Name:          g.Name,
		Enabled:       functions.GetVal(g.Enabled),
		IPAddressType: functions.GetVal(g.IpAddressType),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: g.AWSCloudFormationMetadata,
	}}

}
