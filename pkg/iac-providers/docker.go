

package iacprovider

import (
	"reflect"

	dockerv1 "github.com/khulnasoft/terrasec/pkg/iac-providers/docker/v1"
)

// docker specific constants
const (
	docker                  supportedIacType    = "docker"
	dockerV1                supportedIacVersion = "v1"
	dockerDefaultIacVersion                     = dockerV1
)

// register docker as an IaC provider with terrasec
func init() {
	// register iac provider
	RegisterIacProvider(docker, dockerV1, dockerDefaultIacVersion, reflect.TypeOf(dockerv1.DockerV1{}))
}
