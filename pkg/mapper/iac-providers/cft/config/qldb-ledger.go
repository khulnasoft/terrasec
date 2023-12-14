

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/qldb"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// QldbLedgerConfig holds config for aws_qldb_ledger resource
type QldbLedgerConfig struct {
	Config
	Name               string `json:"name,omitempty"`
	PermissionsMode    string `json:"permissions_mode"`
	DeletionProtection bool   `json:"deletion_protection"`
}

// GetQldbLedgerConfig returns config for aws_qldb_ledger resource
func GetQldbLedgerConfig(q *qldb.Ledger) []AWSResourceConfig {

	cf := QldbLedgerConfig{
		Config: Config{
			Name: functions.GetVal(q.Name),
			Tags: q.Tags,
		},
		Name:               functions.GetVal(q.Name),
		PermissionsMode:    q.PermissionsMode,
		DeletionProtection: functions.GetVal(q.DeletionProtection),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: q.AWSCloudFormationMetadata,
	}}
}
