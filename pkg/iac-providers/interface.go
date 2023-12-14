package iacprovider

import (
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
)

// IacProvider defines the interface which every IaC provider needs to implement
// to claim support in terrasec
type IacProvider interface {
	LoadIacFile(string, map[string]interface{}) (output.AllResourceConfigs, error)
	LoadIacDir(string, map[string]interface{}) (output.AllResourceConfigs, error)
	Name() string
}
