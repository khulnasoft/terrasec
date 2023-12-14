

package policy

const (
	gcp                  supportedCloudType  = "gcp"
	defaultGCPIacType    supportedIacType    = "terraform"
	defaultGCPIacVersion supportedIacVersion = version12
)

func init() {
	// Register gcp as a cloud provider with terrasec
	RegisterCloudProvider(gcp, defaultGCPIacType, defaultGCPIacVersion)
}
