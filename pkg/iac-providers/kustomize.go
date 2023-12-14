package iacprovider

import (
	"reflect"

	kustomizev2 "github.com/khulnasoft/terrasec/pkg/iac-providers/kustomize/v2"
	kustomizev3 "github.com/khulnasoft/terrasec/pkg/iac-providers/kustomize/v3"
	kustomizev4 "github.com/khulnasoft/terrasec/pkg/iac-providers/kustomize/v4"
)

// kustomize specific constants
const (
	kustomize                  supportedIacType    = "kustomize"
	kustomizeV4                supportedIacVersion = "v4"
	kustomizeV3                supportedIacVersion = "v3"
	kustomizeV2                supportedIacVersion = "v2"
	kustomizeDefaultIacVersion                     = kustomizeV4
)

// register kustomize as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(kustomize, kustomizeV4, kustomizeDefaultIacVersion, reflect.TypeOf(kustomizev4.KustomizeV4{}))
	RegisterIacProvider(kustomize, kustomizeV3, kustomizeDefaultIacVersion, reflect.TypeOf(kustomizev3.KustomizeV3{}))
	RegisterIacProvider(kustomize, kustomizeV2, kustomizeDefaultIacVersion, reflect.TypeOf(kustomizev2.KustomizeV2{}))
}
