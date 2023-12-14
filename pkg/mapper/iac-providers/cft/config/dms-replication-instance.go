package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/dms"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// DmsReplicationInstanceConfig holds config for DmsReplicationInstance
type DmsReplicationInstanceConfig struct {
	Config
	AllocatedStorage           int      `json:"allocated_storage"`
	AutoMinorVersionUpgrade    bool     `jons:"auto_minor_version_upgrade"`
	AvailabilityZone           string   `json:"availability_zone"`
	EngineVersion              string   `json:"engine_version"`
	KMSKeyARN                  string   `json:"kms_key_arn"`
	MultiAZ                    bool     `json:"multi_az"`
	PreferredMaintenanceWindow string   `json:"preferred_maintenance_window"`
	PubliclyAccessible         bool     `json:"publicly_accessible"`
	ReplicationInstanceClass   string   `json:"replication_instance_class"`
	ReplicationInstanceID      string   `json:"replication_instance_id"`
	ReplicationSubnetGroupID   string   `json:"replication_subnet_group_id"`
	VPCSecurityGroupIDs        []string `json:"vpc_security_group_ids"`
}

// GetDmsReplicationInstanceConfig returns config for DmsReplicationInstance
func GetDmsReplicationInstanceConfig(r *dms.ReplicationInstance) []AWSResourceConfig {
	cf := DmsReplicationInstanceConfig{
		Config: Config{
			Tags: r.Tags,
		},

		AllocatedStorage:           functions.GetVal(r.AllocatedStorage),
		AutoMinorVersionUpgrade:    functions.GetVal(r.AutoMinorVersionUpgrade),
		AvailabilityZone:           functions.GetVal(r.AvailabilityZone),
		EngineVersion:              functions.GetVal(r.EngineVersion),
		KMSKeyARN:                  functions.GetVal(r.KmsKeyId),
		MultiAZ:                    functions.GetVal(r.MultiAZ),
		PreferredMaintenanceWindow: functions.GetVal(r.PreferredMaintenanceWindow),
		PubliclyAccessible:         functions.GetVal(r.PubliclyAccessible),
		ReplicationInstanceClass:   r.ReplicationInstanceClass,
		ReplicationInstanceID:      functions.GetVal(r.ReplicationInstanceIdentifier),
		ReplicationSubnetGroupID:   functions.GetVal(r.ReplicationSubnetGroupIdentifier),
		VPCSecurityGroupIDs:        r.VpcSecurityGroupIds,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
