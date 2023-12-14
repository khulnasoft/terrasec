

package armv1

import "github.com/hashicorp/go-multierror"

// ARMV1 struct implements the IacProvider interface
type ARMV1 struct {
	templateParameters map[string]interface{}

	errIacLoadDirs *multierror.Error
	// absRootDir is the root directory being scanned.
	// if a file scan was initiated, absRootDir should be empty.
	absRootDir string
}

const (
	// JSONExtension json
	JSONExtension = ".json"

	// ParametersFileExtension .parameters.json
	ParametersFileExtension = ".parameters.json"

	// MetadataFileExtension .metadata.json
	MetadataFileExtension = ".metadata.json"

	// UnknownExtension unknown
	UnknownExtension = "unknown"
)

// ARMFileExtensions returns the valid extensions for Azure ARM (json)
func ARMFileExtensions() []string {
	return []string{JSONExtension}
}
