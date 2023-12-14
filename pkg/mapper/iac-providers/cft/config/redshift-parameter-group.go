

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/redshift"
)

// ParameterBlock holds config for Parameter
type ParameterBlock struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RedshiftParameterGroupConfig holds config for RedshiftParameterGroup
type RedshiftParameterGroupConfig struct {
	Config
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Family      string           `json:"family"`
	Parameter   []ParameterBlock `json:"parameter"`
}

// GetRedshiftParameterGroupConfig returns config for RedshiftParameterGroup
func GetRedshiftParameterGroupConfig(p *redshift.ClusterParameterGroup, paramGroupName string) []AWSResourceConfig {
	var parameterBlock []ParameterBlock
	if p.Parameters != nil {
		parameterBlock := make([]ParameterBlock, len(p.Parameters))
		for i, parameter := range p.Parameters {
			parameterBlock[i].Name = parameter.ParameterName
			parameterBlock[i].Value = parameter.ParameterValue
		}
	}

	cf := RedshiftParameterGroupConfig{
		Config: Config{
			Name: paramGroupName,
			Tags: p.Tags,
		},
		Name:        paramGroupName,
		Description: p.Description,
		Family:      p.ParameterGroupFamily,
		Parameter:   parameterBlock,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: p.AWSCloudFormationMetadata,
	}}
}
