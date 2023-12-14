

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/elasticloadbalancingv2"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// ElasticLoadBalancingV2ListenerConfig holds config for aws_lb_listener
type ElasticLoadBalancingV2ListenerConfig struct {
	Config
	Protocol      string                `json:"protocol"`
	Port          int                   `json:"port"`
	DefaultAction []DefaultActionConfig `json:"default_action"`
}

// DefaultActionConfig holds config for default_action attribute of aws_lb_listener
type DefaultActionConfig struct {
	RedirectConfig []RedirectConfig `json:"redirect"`
}

// RedirectConfig holds config for redirect attribute of default_action
type RedirectConfig struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
}

// GetElasticLoadBalancingV2ListenerConfig returns config for aws_lb_listener
func GetElasticLoadBalancingV2ListenerConfig(l *elasticloadbalancingv2.Listener) []AWSResourceConfig {
	// create a listener subresource per DefaultAction defined in cft
	// as only one default action per listener is possible in terraform
	resourceConfigs := make([]AWSResourceConfig, 0)

	for _, action := range l.DefaultActions {
		// DefaultActions are required
		cf := ElasticLoadBalancingV2ListenerConfig{
			Config:   Config{},
			Protocol: functions.GetVal(l.Protocol),
			Port:     functions.GetVal(l.Port),
		}
		if action.RedirectConfig != nil {
			defaultAction := []DefaultActionConfig{
				{
					RedirectConfig: []RedirectConfig{
						{
							Protocol: functions.GetVal(action.RedirectConfig.Protocol),
							Port:     functions.GetVal(action.RedirectConfig.Port),
						},
					},
				},
			}
			cf.DefaultAction = defaultAction
		}
		resourceConfigs = append(resourceConfigs, AWSResourceConfig{
			Resource: cf,
			Metadata: l.AWSCloudFormationMetadata,
		})
	}

	return resourceConfigs
}
