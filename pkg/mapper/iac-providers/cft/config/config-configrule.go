package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/config"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// AWSConfigConfigRuleConfig holds config for aws_config_config_rule
type AWSConfigConfigRuleConfig struct {
	Config
	Source interface{} `json:"source"`
}

// GetConfigConfigRuleConfig returns config for aws_config_config_rule
func GetConfigConfigRuleConfig(c *config.ConfigRule) []AWSResourceConfig {
	cf := AWSConfigConfigRuleConfig{
		Config: Config{
			Name: functions.GetVal(c.ConfigRuleName),
		},
	}
	if c.Source != nil {
		sources := make([]map[string]interface{}, 0)
		source := make(map[string]interface{})
		source["source_identifier"] = functions.GetVal(c.Source.SourceIdentifier)
		sources = append(sources, source)
		if len(sources) > 0 {
			cf.Source = sources
		}
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
