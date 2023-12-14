

package dockerv1

import "github.com/hashicorp/go-multierror"

// DockerV1 struct implements the IacProvider interface
type DockerV1 struct {
	errIacLoadDirs *multierror.Error
	// absRootDir is the root directory being scanned.
	// if a file scan was initiated, absRootDir should be empty.
	absRootDir string
}

const (
	// DockerFileName dockerfile name to be used when directory path is given
	DockerFileName = "Dockerfile"
)
