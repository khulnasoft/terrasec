

package core

import "github.com/khulnasoft/terrasec/pkg/iac-providers/output"

// Mapper defines the base API that each IaC provider mapper must implement.
type Mapper interface {
	// Map transforms the provider specific template to terrasec native format.
	Map(resource interface{}, params ...map[string]interface{}) ([]output.ResourceConfig, error)
}
