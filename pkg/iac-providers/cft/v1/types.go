

package cftv1

import "github.com/hashicorp/go-multierror"

// CFTV1 struct implements the IacProvider interface
type CFTV1 struct {
	errIacLoadDirs *multierror.Error
	// absRootDir is the root directory being scanned.
	// if a file scan was initiated, absRootDir should be empty.
	absRootDir string
}

const (
	// YAMLExtension yaml
	YAMLExtension = "yaml"

	// YAMLExtension2 yml
	YAMLExtension2 = "yml"

	// JSONExtension json
	JSONExtension = "json"

	// TXTExtension txt
	TXTExtension = "txt"

	// TemplateExtension template
	TemplateExtension = "template"

	// UnknownExtension unknown
	UnknownExtension = "unknown"
)

// CFTFileExtensions returns the valid extensions for AWS CFT (json | YAML | txt | template)
func CFTFileExtensions() []string {
	return []string{YAMLExtension, YAMLExtension2, JSONExtension, TemplateExtension, TXTExtension}
}

type cftResource struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion"`
	Resources                map[string]interface{} `json:"Resources"`
}
