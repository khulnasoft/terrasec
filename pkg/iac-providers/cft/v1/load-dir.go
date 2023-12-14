package cftv1

import (
	"fmt"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/hashicorp/go-multierror"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/results"
	"github.com/khulnasoft/terrasec/pkg/utils"
)

// LoadIacDir loads all CFT template files in the current directory.
func (a *CFTV1) LoadIacDir(absRootDir string, options map[string]interface{}) (output.AllResourceConfigs, error) {
	a.absRootDir = absRootDir

	allResourcesConfig := make(map[string][]output.ResourceConfig)

	cftFileMap, err := utils.FindFilesBySuffix(absRootDir, CFTFileExtensions())
	if err != nil {
		zap.S().Debug("error while searching for iac files", zap.String("root dir", absRootDir), zap.Error(err))
		return allResourcesConfig, multierror.Append(a.errIacLoadDirs, results.DirScanErr{IacType: "cft", Directory: absRootDir, ErrMessage: err.Error()})
	}

	if len(cftFileMap) == 0 {
		errMsg := fmt.Sprintf("cft files not found in the directory %s", a.absRootDir)
		return allResourcesConfig, multierror.Append(a.errIacLoadDirs, results.DirScanErr{IacType: "cft", Directory: a.absRootDir, ErrMessage: errMsg})
	}

	for fileDir, files := range cftFileMap {
		for i := range files {
			file := filepath.Join(fileDir, *files[i])

			var configData output.AllResourceConfigs
			if configData, err = a.LoadIacFile(file, options); err != nil {
				errMsg := fmt.Sprintf("error while loading iac file '%s', err: %v", file, err)
				zap.S().Debug("error while loading iac files", zap.String("IAC file", file), zap.Error(err))
				a.errIacLoadDirs = multierror.Append(a.errIacLoadDirs, results.DirScanErr{IacType: "cft", Directory: fileDir, ErrMessage: errMsg})
				continue
			}

			for key := range configData {
				allResourcesConfig[key] = append(allResourcesConfig[key], configData[key]...)
			}
		}
	}

	return allResourcesConfig, a.errIacLoadDirs
}

// Name returns name of the provider
func (*CFTV1) Name() string {
	return "cft"
}
