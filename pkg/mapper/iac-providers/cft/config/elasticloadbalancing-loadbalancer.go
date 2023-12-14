

package config

import (
	"fmt"

	"github.com/awslabs/goformation/v7/cloudformation/elasticloadbalancing"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// GetPolicies represents subresource aws_load_balancer_policy for Policies attribute
const (
	GetPolicies = "Policies"
)

// PolicyAttributeBlock holds config for PolicyTypeBlock
type PolicyAttributeBlock struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ElasticLoadBalancingLoadBalancerPoliciesConfig holds config for ElasticLoadBalancingLoadBalancerPolicies
type ElasticLoadBalancingLoadBalancerPoliciesConfig struct {
	Config
	LoadBalancerName string                 `json:"load_balancer_name"`
	PolicyName       string                 `json:"policy_name"`
	PolicyTypeName   string                 `jons:"policy_type_name"`
	PolicyAttribute  []PolicyAttributeBlock `json:"policy_attribute"`
}

// ElasticLoadBalancingLoadBalancerConfig holds config for aws_elb
type ElasticLoadBalancingLoadBalancerConfig struct {
	Listeners           interface{} `json:"listener"`
	AccessLoggingPolicy interface{} `json:"access_logs,omitempty"`
	Config
}

// ELBAccessLoggingPolicyConfig holds config for access_logs attribute of aws_elb
type ELBAccessLoggingPolicyConfig struct {
	Enabled bool `json:"enabled"`
}

// ELBListenerConfig holds config for listener attribute of aws_elb
type ELBListenerConfig struct {
	LBProtocol       string `json:"lb_protocol"`
	InstanceProtocol string `json:"instance_protocol"`
}

// GetElasticLoadBalancingLoadBalancerConfig returns config for aws_elb
func GetElasticLoadBalancingLoadBalancerConfig(e *elasticloadbalancing.LoadBalancer, elbname string) []AWSResourceConfig {
	elbpolicies := make([]ElasticLoadBalancingLoadBalancerPoliciesConfig, len(e.Policies))
	awsconfig := make([]AWSResourceConfig, len(e.Policies))

	for i, policy := range e.Policies {
		indexedElbName := fmt.Sprintf("%s%d", elbname, i)

		elbpolicies[i].LoadBalancerName = indexedElbName
		elbpolicies[i].PolicyName = policy.PolicyName
		elbpolicies[i].PolicyTypeName = policy.PolicyType

		elbpolicies[i].PolicyAttribute = make([]PolicyAttributeBlock, len(policy.Attributes))
		for ai := range policy.Attributes {
			attribVals, ok := policy.Attributes[ai].(map[string]interface{})
			if !ok {
				continue
			}

			elbpolicies[i].PolicyAttribute[ai].Name, ok = attribVals["Name"].(string)
			if !ok {
				continue
			}

			elbpolicies[i].PolicyAttribute[ai].Value, ok = attribVals["Value"].(string)
			if !ok {
				continue
			}

			// variable "ok" is only used for safe type conversion
			_ = ok
		}

		awsconfig[i].Type = GetPolicies
		awsconfig[i].Name = indexedElbName
		awsconfig[i].Resource = elbpolicies[i]
		awsconfig[i].Metadata = e.AWSCloudFormationMetadata
	}

	cf := ElasticLoadBalancingLoadBalancerConfig{
		Config: Config{
			Tags: e.Tags,
		},
	}

	if e.AccessLoggingPolicy != nil {
		cf.AccessLoggingPolicy = ELBAccessLoggingPolicyConfig{
			Enabled: e.AccessLoggingPolicy.Enabled,
		}
	}

	if e.Listeners != nil {
		lc := make([]ELBListenerConfig, 0)
		for _, listener := range e.Listeners {
			lc = append(lc, ELBListenerConfig{
				InstanceProtocol: functions.GetVal(listener.InstanceProtocol),
				LBProtocol:       listener.Protocol,
			})
		}
		cf.Listeners = lc
	}

	var awsconfigElb AWSResourceConfig
	awsconfigElb.Resource = cf
	awsconfigElb.Metadata = e.AWSCloudFormationMetadata
	awsconfig = append(awsconfig, awsconfigElb)

	return awsconfig
}
