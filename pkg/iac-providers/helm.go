package iacprovider

import (
	"reflect"

	helmv3 "github.com/khulnasoft/terrasec/pkg/iac-providers/helm/v3"
)

// terraform specific constants
const (
	helm                  supportedIacType    = "helm"
	helmV3                supportedIacVersion = "v3"
	helmDefaultIacVersion                     = helmV3
)

// register helm as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(helm, helmV3, helmDefaultIacVersion, reflect.TypeOf(helmv3.HelmV3{}))
}
