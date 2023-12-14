package policy

const (
	defaultAllIacType    supportedIacType    = "terraform"
	defaultAllIacVersion supportedIacVersion = version12
)

func init() {
	// Register all as a cloud provider with terrasec
	RegisterIndirectCloudProvider("all", defaultAllIacType, defaultAllIacVersion, func() []string {
		return SupportedPolicyTypes(false)
	})
}
