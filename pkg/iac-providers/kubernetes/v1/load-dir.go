

package k8sv1

import (
	"fmt"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"github.com/hashicorp/go-multierror"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/results"
	"github.com/khulnasoft/terrasec/pkg/utils"
)

func (*K8sV1) getFileType(file string) string {
	if strings.HasSuffix(file, YAMLExtension) {
		return YAMLExtension
	} else if strings.HasSuffix(file, YAMLExtension2) {
		return YAMLExtension2
	} else if strings.HasSuffix(file, JSONExtension) {
		return JSONExtension
	}
	return UnknownExtension
}

// LoadIacDir loads all k8s files in the current directory
func (k *K8sV1) LoadIacDir(absRootDir string, options map[string]interface{}) (output.AllResourceConfigs, error) {
	// set the root directory being scanned
	k.absRootDir = absRootDir

	allResourcesConfig := make(map[string][]output.ResourceConfig)

	fileMap, err := utils.FindFilesBySuffix(absRootDir, K8sFileExtensions())
	if err != nil {
		zap.S().Debug("error while searching for iac files", zap.String("root dir", absRootDir), zap.Error(err))
		return allResourcesConfig, multierror.Append(k.errIacLoadDirs, results.DirScanErr{IacType: "k8s", Directory: absRootDir, ErrMessage: err.Error()})
	}
	if len(fileMap) == 0 {
		errMsg := fmt.Sprintf("kubernetes files not found in the directory %s", k.absRootDir)
		return allResourcesConfig, multierror.Append(k.errIacLoadDirs, results.DirScanErr{IacType: "k8s", Directory: k.absRootDir, ErrMessage: errMsg})
	}

	for fileDir, files := range fileMap {
		for i := range files {
			file := filepath.Join(fileDir, *files[i])

			var configData output.AllResourceConfigs
			if configData, err = k.LoadIacFile(file, options); err != nil {
				errMsg := fmt.Sprintf("error while loading iac file '%s'. err: %v", file, err)
				zap.S().Debug("error while loading iac files", zap.String("IAC file", file), zap.Error(err))
				k.errIacLoadDirs = multierror.Append(k.errIacLoadDirs, results.DirScanErr{IacType: "k8s", Directory: fileDir, ErrMessage: errMsg})
				continue
			}

			for key := range configData {
				allResourcesConfig[key] = append(allResourcesConfig[key], configData[key]...)
			}
		}
	}

	return allResourcesConfig, k.errIacLoadDirs
}

// Name returns name of the provider
func (*K8sV1) Name() string {
	return kubernetesTypeNameShort
}
