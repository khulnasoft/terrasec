

package k8sv1

import "github.com/hashicorp/go-multierror"

// K8sV1 struct implements the IacProvider interface
type K8sV1 struct {
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

	// UnknownExtension unknown
	UnknownExtension = "unknown"

	kubernetesTypeName      = "kubernetes"
	defaultNamespace        = "default"
	kubernetesTypeNameShort = "k8s"
)

// K8sFileExtensions returns the valid extensions for k8s (yaml, yml, json)
func K8sFileExtensions() []string {
	return []string{YAMLExtension, YAMLExtension2, JSONExtension}
}
