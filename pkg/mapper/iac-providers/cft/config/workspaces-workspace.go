

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/workspaces"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// WorkspacesWorkspaceConfig holds config for aws_workspaces_workspace
type WorkspacesWorkspaceConfig struct {
	Config
	RootVolumeEncryptionEnabled bool `json:"root_volume_encryption_enabled,omitempty"`
	UserVolumeEncryptionEnabled bool `json:"user_volume_encryption_enabled,omitempty"`
}

// GetWorkspacesWorkspaceConfig returns config for aws_workspaces_workspace
func GetWorkspacesWorkspaceConfig(w *workspaces.Workspace) []AWSResourceConfig {
	cf := WorkspacesWorkspaceConfig{
		Config: Config{
			Name: w.UserName,
			Tags: w.Tags,
		},
		UserVolumeEncryptionEnabled: functions.GetVal(w.UserVolumeEncryptionEnabled),
		RootVolumeEncryptionEnabled: functions.GetVal(w.RootVolumeEncryptionEnabled),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: w.AWSCloudFormationMetadata,
	}}
}
