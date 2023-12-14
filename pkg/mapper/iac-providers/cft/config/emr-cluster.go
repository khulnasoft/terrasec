package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/emr"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// KerberosAttributesBlock holds config for KerberosAttributes
type KerberosAttributesBlock struct {
	KDCAdminPassword string `json:"kdc_admin_password"`
	Realm            string `json:"realm"`
}

// EmrClusterConfig holds config for EmrCluster
type EmrClusterConfig struct {
	Config
	Name                  string                    `json:"name"`
	ReleaseLabel          string                    `json:"release_label"`
	ServiceRole           string                    `json:"service_role"`
	TerminationProtection bool                      `json:"termination_protection"`
	KerberosAttributes    []KerberosAttributesBlock `json:"kerberos_attributes"`
}

// GetEmrClusterConfig returns config for EmrCluster
func GetEmrClusterConfig(c *emr.Cluster) []AWSResourceConfig {
	var kerberosAttributes []KerberosAttributesBlock
	if c.KerberosAttributes != nil {
		kerberosAttributes = make([]KerberosAttributesBlock, 1)

		kerberosAttributes[0].KDCAdminPassword = c.KerberosAttributes.KdcAdminPassword
		kerberosAttributes[0].Realm = c.KerberosAttributes.Realm
	}

	cf := EmrClusterConfig{
		Config: Config{
			Name: c.Name,
		},
		Name:                  c.Name,
		ReleaseLabel:          functions.GetVal(c.ReleaseLabel),
		ServiceRole:           c.ServiceRole,
		TerminationProtection: functions.GetVal(c.Instances.TerminationProtected),
		KerberosAttributes:    kerberosAttributes,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
