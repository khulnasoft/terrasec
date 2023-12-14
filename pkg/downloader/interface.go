

package downloader

import (
	hclConfigs "github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/registry/regsrc"
	"github.com/hashicorp/terraform/registry/response"
)

// Downloader helps in downloading different kinds of modules from
// different types of sources
type Downloader interface {
	Download(url, destDir string) (finalDir string, err error)
	DownloadWithType(remoteType, url, dest string) (finalDir string, err error)
	GetURLSubDir(url, dest string) (urlWithType string, subDir string, err error)
	SubDirGlob(string, string) (string, error)
}

// ModuleDownloader helps in downloading the remote modules
type ModuleDownloader interface {
	DownloadModule(addr, destPath string) (string, error)
	DownloadRemoteModule(requiredVersion hclConfigs.VersionConstraint, destPath string, module *regsrc.Module) (string, error)
	CleanUp()
	GetDownloaderCache() map[string]string
}

// terraformRegistryClient will help interact with terraform registries
type terraformRegistryClient interface {
	ModuleVersions(module *regsrc.Module) (*response.ModuleVersions, error)
	ModuleLocation(module *regsrc.Module, version string) (string, error)
}

// NewDownloader returns a new downloader
func NewDownloader() Downloader {
	return newGoGetter()
}

// NewRemoteDownloader returns a new ModuleDownloader
func NewRemoteDownloader() ModuleDownloader {
	return newRemoteModuleInstaller()
}

// newClientRegistry returns a terraformClientRegistry to query terraform registries
func newClientRegistry() terraformRegistryClient {
	return newTerraformRegistryClient()
}
