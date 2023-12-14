package iacprovider

import (
	"reflect"

	k8sv1 "github.com/khulnasoft/terrasec/pkg/iac-providers/kubernetes/v1"
)

// terraform specific constants
const (
	kubernetes                  supportedIacType    = "k8s"
	kubernetesV1                supportedIacVersion = "v1"
	kubernetesDefaultIacVersion                     = kubernetesV1
)

// register kubernetes as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(kubernetes, kubernetesV1, kubernetesDefaultIacVersion, reflect.TypeOf(k8sv1.K8sV1{}))
}
