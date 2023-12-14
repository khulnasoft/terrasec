

package iacprovider

import (
	"reflect"

	armv1 "github.com/khulnasoft/terrasec/pkg/iac-providers/arm/v1"
)

// terraform specific constants
const (
	arm                  supportedIacType    = "arm"
	armV1                supportedIacVersion = "v1"
	armDefaultIacVersion                     = armV1
)

// register arm as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(arm, armV1, armDefaultIacVersion, reflect.TypeOf(armv1.ARMV1{}))
}
