

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/msk"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// EncryptionInTransitBlock holds config for EncryptionInTransit
type EncryptionInTransitBlock struct {
	ClientBroker string `json:"client_broker"`
	InCluster    bool   `json:"in_cluster"`
}

// EncryptionInfoBlock holds config for EncryptionInfo
type EncryptionInfoBlock struct {
	EncryptionAtRestKmsKeyArn string                     `json:"encryption_at_rest_kms_key_arn"`
	EncryptionInTransit       []EncryptionInTransitBlock `json:"encryption_in_transit"`
}

// BrokerNodeGroupInfoBlock holds config for BrokerNodeGroupInfo
type BrokerNodeGroupInfoBlock struct {
	InstanceType   string   `json:"instance_type"`
	EksVolumeSize  int      `json:"eks_volume_size"`
	ClientSubnets  []string `json:"client_subnets"`
	SecurityGroups []string `json:"security_groups"`
}

// MskClusterConfig holds config for MskCluster
type MskClusterConfig struct {
	Config
	ClusterName         string                     `json:"cluster_name"`
	KafkaVersion        string                     `json:"kafka_version"`
	NumberOfBrokerNodes int                        `json:"number_of_broker_nodes"`
	BrokerNodeGroupInfo []BrokerNodeGroupInfoBlock `json:"broker_node_group_info"`
	EncryptionInfo      []EncryptionInfoBlock      `json:"encryption_info"`
}

// GetMskClusterConfig returns config for MskCluster
func GetMskClusterConfig(c *msk.Cluster) []AWSResourceConfig {
	var brokerNodeGroupInfo []BrokerNodeGroupInfoBlock
	if c.BrokerNodeGroupInfo != nil {
		brokerNodeGroupInfo = make([]BrokerNodeGroupInfoBlock, 1)

		brokerNodeGroupInfo[0].InstanceType = c.BrokerNodeGroupInfo.InstanceType
		brokerNodeGroupInfo[0].EksVolumeSize = functions.GetVal(c.BrokerNodeGroupInfo.StorageInfo.EBSStorageInfo.VolumeSize)
		brokerNodeGroupInfo[0].ClientSubnets = c.BrokerNodeGroupInfo.ClientSubnets
		brokerNodeGroupInfo[0].SecurityGroups = c.BrokerNodeGroupInfo.SecurityGroups
	}

	var encryptionInfo []EncryptionInfoBlock
	if c.EncryptionInfo != nil {
		encryptionInfo = make([]EncryptionInfoBlock, 1)

		if c.EncryptionInfo.EncryptionAtRest != nil {
			encryptionInfo[0].EncryptionAtRestKmsKeyArn = c.EncryptionInfo.EncryptionAtRest.DataVolumeKMSKeyId
		}

		if c.EncryptionInfo.EncryptionInTransit != nil {
			encryptionInfo[0].EncryptionInTransit = make([]EncryptionInTransitBlock, 1)

			encryptionInfo[0].EncryptionInTransit[0].ClientBroker = functions.GetVal(c.EncryptionInfo.EncryptionInTransit.ClientBroker)
			encryptionInfo[0].EncryptionInTransit[0].InCluster = functions.GetVal(c.EncryptionInfo.EncryptionInTransit.InCluster)
		}
	}

	cf := MskClusterConfig{
		Config: Config{
			Name: c.ClusterName,
			Tags: c.Tags,
		},
		ClusterName:         c.ClusterName,
		KafkaVersion:        c.KafkaVersion,
		NumberOfBrokerNodes: c.NumberOfBrokerNodes,
		BrokerNodeGroupInfo: brokerNodeGroupInfo,
		EncryptionInfo:      encryptionInfo,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
