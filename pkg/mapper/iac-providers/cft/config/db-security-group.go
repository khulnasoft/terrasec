package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/rds"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// DBIngress holds config for ingress block
type DBIngress struct {
	CIDR              string `json:"cidr"`
	SecurityGroupName string `json:"security_group_name"`
}

// DBSecurityGroupConfig holds config for aws_db_security_group
type DBSecurityGroupConfig struct {
	Config
	Ingress []DBIngress `json:"ingress"`
}

// GetDBSecurityGroupConfig returns config for aws_db_security_group
func GetDBSecurityGroupConfig(dbsg *rds.DBSecurityGroup) []AWSResourceConfig {
	cf := DBSecurityGroupConfig{
		Config: Config{
			Tags: dbsg.Tags,
		},
	}

	if dbsg.DBSecurityGroupIngress != nil {
		cf.Ingress = make([]DBIngress, len(dbsg.DBSecurityGroupIngress))
		for i, dbsgi := range dbsg.DBSecurityGroupIngress {
			cf.Ingress[i].CIDR = functions.GetVal(dbsgi.CIDRIP)
			cf.Ingress[i].SecurityGroupName = functions.GetVal(dbsgi.EC2SecurityGroupName)
		}
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: dbsg.AWSCloudFormationMetadata,
	}}
}
