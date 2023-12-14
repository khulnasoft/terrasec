

package rego

import (
	"encoding/json"
	"os"

	"github.com/khulnasoft/terrasec/pkg/policy"
)

// LoadRegoMetadata reads rego meta data file
func LoadRegoMetadata(file string) (*policy.RegoMetadata, error) {
	metadata, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// Read metadata into struct
	regoMetadata := policy.RegoMetadata{}
	if err = json.Unmarshal(metadata, &regoMetadata); err != nil {
		return nil, err
	}
	return &regoMetadata, nil
}
