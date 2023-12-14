

package policy

const (
	helm                  supportedCloudType  = "k8s"
	defaultHelmIacType    supportedIacType    = "helm"
	defaultHelmIacVersion supportedIacVersion = version3
)

func init() {
	// Register helm as a provider with terrasec
	RegisterCloudProvider(helm, defaultHelmIacType, defaultHelmIacVersion)
}
