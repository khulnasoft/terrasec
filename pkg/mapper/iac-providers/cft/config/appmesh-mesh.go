package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/appmesh"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// AppMeshEgressFilterBlock holds config for AppMeshEgressFilter
type AppMeshEgressFilterBlock struct {
	Type string `json:"type"`
}

// AppMeshSpecBlock holds config for AppMeshSpec
type AppMeshSpecBlock struct {
	EgressFilter []AppMeshEgressFilterBlock `json:"egress_filter"`
}

// AppMeshMeshConfig holds config for AppMeshMesh
type AppMeshMeshConfig struct {
	Config
	Name string             `json:"name"`
	Spec []AppMeshSpecBlock `json:"spec"`
}

// GetAppMeshMeshConfig returns config for AppMeshMesh
func GetAppMeshMeshConfig(m *appmesh.Mesh) []AWSResourceConfig {
	var spec []AppMeshSpecBlock
	if m.Spec != nil {
		spec = make([]AppMeshSpecBlock, 1)

		if m.Spec.EgressFilter != nil {
			spec[0].EgressFilter = make([]AppMeshEgressFilterBlock, 1)

			spec[0].EgressFilter[0].Type = m.Spec.EgressFilter.Type
		}
	}

	cf := AppMeshMeshConfig{
		Config: Config{
			Name: functions.GetVal(m.MeshName),
			Tags: m.Tags,
		},
		Name: functions.GetVal(m.MeshName),
		Spec: spec,
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: m.AWSCloudFormationMetadata,
	}}
}
