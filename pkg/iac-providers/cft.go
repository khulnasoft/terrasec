package iacprovider

import (
	"reflect"

	cftv1 "github.com/khulnasoft/terrasec/pkg/iac-providers/cft/v1"
)

// terraform specific constants
const (
	cft                  supportedIacType    = "cft"
	cftV1                supportedIacVersion = "v1"
	cftDefaultIacVersion                     = cftV1
)

// register cft as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(cft, cftV1, cftDefaultIacVersion, reflect.TypeOf(cftv1.CFTV1{}))
}
