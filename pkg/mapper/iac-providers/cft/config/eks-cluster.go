

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/eks"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// EKSVPCConfigBlock holds config for EKSVPCConfig
type EKSVPCConfigBlock struct {
	SubnetIDs             []string `json:"subnet_ids"`
	SecurityGroupIDs      []string `json:"security_group_ids"`
	EndpointPrivateAccess bool     `json:"endpoint_private_access"`
	EndpointPublicAccess  bool     `json:"endpoint_public_access"`
}

// EksClusterConfig holds config for EksCluster
type EksClusterConfig struct {
	Config
	Name                   string              `json:"name"`
	RoleARN                string              `json:"role_arn"`
	VPCConfig              []EKSVPCConfigBlock `json:"vpc_config"`
	EnabledClusterLogTypes []string            `json:"enabled_cluster_log_types"`
}

// GetEksClusterConfig returns config for EksCluster
func GetEksClusterConfig(c *eks.Cluster) []AWSResourceConfig {
	var vpcConfig []EKSVPCConfigBlock
	if c.ResourcesVpcConfig != nil {
		vpcConfig := make([]EKSVPCConfigBlock, 1)

		vpcConfig[0].SubnetIDs = c.ResourcesVpcConfig.SubnetIds
		vpcConfig[0].SecurityGroupIDs = c.ResourcesVpcConfig.SecurityGroupIds
		vpcConfig[0].EndpointPrivateAccess = functions.GetVal(c.ResourcesVpcConfig.EndpointPrivateAccess)
		vpcConfig[0].EndpointPublicAccess = functions.GetVal(c.ResourcesVpcConfig.EndpointPublicAccess)
	}

	cf := EksClusterConfig{
		Config: Config{
			Name: functions.GetVal(c.Name),
		},
		Name:      functions.GetVal(c.Name),
		RoleARN:   c.RoleArn,
		VPCConfig: vpcConfig,
	}

	enabledTypes := c.Logging.ClusterLogging.EnabledTypes
	if c.Logging != nil && c.Logging.ClusterLogging != nil && len(enabledTypes) > 0 {
		enabledClusterLogTypes := make([]string, len(enabledTypes))
		for i, enabledType := range enabledTypes {
			enabledClusterLogTypes[i] = functions.GetVal(enabledType.Type)
		}
		cf.EnabledClusterLogTypes = enabledClusterLogTypes
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
