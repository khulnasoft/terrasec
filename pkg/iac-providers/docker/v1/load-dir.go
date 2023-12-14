package dockerv1

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/results"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
)

// LoadIacDir loads the docker file specified in given folder.
func (dc *DockerV1) LoadIacDir(absRootDir string, options map[string]interface{}) (output.AllResourceConfigs, error) {
	// set the root directory being scanned
	dc.absRootDir = absRootDir

	allResourcesConfig := make(map[string][]output.ResourceConfig)

	// find all the files in the folder with name `Dockerfile`
	fileMap, err := utils.FindFilesBySuffix(absRootDir, []string{DockerFileName})
	if err != nil {
		zap.S().Errorf("error while searching for iac files", zap.String("root dir", absRootDir), zap.Error(err))
		return allResourcesConfig, multierror.Append(dc.errIacLoadDirs, results.DirScanErr{IacType: "docker", Directory: absRootDir, ErrMessage: err.Error()})
	}

	if len(fileMap) == 0 {
		errMsg := fmt.Sprintf("Dockerfile not found in the directory %s", dc.absRootDir)
		return allResourcesConfig, multierror.Append(dc.errIacLoadDirs, results.DirScanErr{IacType: "docker", Directory: dc.absRootDir, ErrMessage: errMsg})
	}

	for fileDir, files := range fileMap {
		for i := range files {
			file := filepath.Join(fileDir, *files[i])
			var configData output.AllResourceConfigs
			if configData, err = dc.LoadIacFile(file, options); err != nil {
				errMsg := fmt.Sprintf("error while parsing file %s", file)
				zap.S().Errorf("error while searching for iac files", zap.String("root dir", absRootDir), errMsg)
				dc.errIacLoadDirs = multierror.Append(dc.errIacLoadDirs, results.DirScanErr{IacType: "docker", Directory: absRootDir, ErrMessage: errMsg})
				continue
			}

			for key := range configData {
				allResourcesConfig[key] = append(allResourcesConfig[key], configData[key]...)
			}
		}
	}

	return allResourcesConfig, dc.errIacLoadDirs

}

// Name returns name of the provider
func (dc *DockerV1) Name() string {
	return "docker"
}
