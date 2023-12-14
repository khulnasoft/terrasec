package policy

const (
	defaultKustomizeIacType    supportedIacType    = "kustomize"
	defaultKustomizeIacVersion supportedIacVersion = version4
)

func init() {
	// Register kustomize as a provider with terrasec
	RegisterCloudProvider(kubernetes, defaultKustomizeIacType, defaultKustomizeIacVersion)
}
