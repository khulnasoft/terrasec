package downloader

import (
	"strings"

	"github.com/hashicorp/terraform/registry/regsrc"
)

const (
	terraformRegistry            = "terraform-registry"
	registryAddrVersionSeparator = ":"
)

var (
	supportedRemoteTypes = []string{"git", "s3", "gcs", "http", terraformRegistry}
	localSourcePrefixes  = []string{
		"./",
		"../",
		".\\",
		"..\\",
	}
)

// IsValidRemoteType validates the remote type supplied as scan option
func IsValidRemoteType(remoteType string) bool {
	for _, r := range supportedRemoteTypes {
		if strings.EqualFold(r, strings.ToLower(strings.TrimSpace(remoteType))) {
			return true
		}
	}
	return false
}

// IsRemoteTypeTerraformRegistry checks if supplied remote type is terraform-registry
func IsRemoteTypeTerraformRegistry(remoteType string) bool {
	return strings.EqualFold(terraformRegistry, strings.ToLower(strings.TrimSpace(remoteType)))
}

// IsLocalSourceAddr validates if a source address is a local address or not
func IsLocalSourceAddr(addr string) bool {
	for _, prefix := range localSourcePrefixes {
		if strings.HasPrefix(addr, prefix) {
			return true
		}
	}
	return false
}

// IsRegistrySourceAddr will validate if the source address is a valid registry
// module or not.
// a valid source address is of the form <HOSTNAME>/NAMESPACE>/<NAME>/<PROVIDER>
// regsrc.ParseModuleSource func returns a terraform registry module source.
func IsRegistrySourceAddr(addr string) bool {
	_, err := regsrc.ParseModuleSource(addr)
	return err == nil
}

// GetSourceAddrAndVersion extracts source address and version from supplied source url
func GetSourceAddrAndVersion(sourceURL string) (string, string) {
	separatorIndex := strings.LastIndex(sourceURL, registryAddrVersionSeparator)
	if separatorIndex == -1 {
		return sourceURL, ""
	}
	return strings.TrimSpace(sourceURL[0:separatorIndex]), strings.TrimSpace(sourceURL[separatorIndex+1:])
}
