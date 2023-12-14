

package policy

const (
	kubernetes                  supportedCloudType  = "k8s"
	defaultKubernetesIacType    supportedIacType    = "k8s"
	defaultKubernetesIacVersion supportedIacVersion = version1
)

func init() {
	// Register kubernetes as a provider with terrasec
	RegisterCloudProvider(kubernetes, defaultKubernetesIacType, defaultKubernetesIacVersion)
}
