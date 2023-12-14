

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ram"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// RAMResourceShareConfig holds config for RAMResourceShare
type RAMResourceShareConfig struct {
	Config
	Name                    string `json:"name"`
	AllowExternalPrincipals bool   `json:"allow_external_principals"`
}

// GetRAMResourceShareConfig returns config for RAMResourceShare
func GetRAMResourceShareConfig(r *ram.ResourceShare) []AWSResourceConfig {
	cf := RAMResourceShareConfig{
		Config: Config{
			Name: r.Name,
			Tags: r.Tags,
		},
		Name:                    r.Name,
		AllowExternalPrincipals: functions.GetVal(r.AllowExternalPrincipals),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
