

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/neptune"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// NeptuneClusterInstanceConfig holds config for aws_neptune_cluster_instance resource
type NeptuneClusterInstanceConfig struct {
	Config
	AutoMinorVersionUpgrade    bool   `json:"auto_minor_version_upgrade,omitempty"`
	AvailabilityZone           string `json:"availability_zone,omitempty"`
	DBClusterIdentifier        string `json:"cluster_identifier,omitempty"`
	DBInstanceClass            string `json:"instance_class,omitempty"`
	DBParameterGroupName       string `json:"neptune_parameter_group_name,omitempty"`
	DBSubnetGroupName          string `json:"neptune_subnet_group_name,omitempty"`
	PreferredMaintenanceWindow string `json:"preferred_backup_window,omitempty"`
}

// GetNeptuneClusterInstanceConfig returns config for aws_neptune_cluster_instance resource
func GetNeptuneClusterInstanceConfig(n *neptune.DBInstance) []AWSResourceConfig {
	cf := NeptuneClusterInstanceConfig{
		Config: Config{
			Tags: n.Tags,
		},
		AutoMinorVersionUpgrade:    functions.GetVal(n.AutoMinorVersionUpgrade),
		AvailabilityZone:           functions.GetVal(n.AvailabilityZone),
		DBClusterIdentifier:        functions.GetVal(n.DBClusterIdentifier),
		DBInstanceClass:            n.DBInstanceClass,
		DBParameterGroupName:       functions.GetVal(n.DBParameterGroupName),
		DBSubnetGroupName:          functions.GetVal(n.DBSubnetGroupName),
		PreferredMaintenanceWindow: functions.GetVal(n.PreferredMaintenanceWindow),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: n.AWSCloudFormationMetadata,
	}}
}
