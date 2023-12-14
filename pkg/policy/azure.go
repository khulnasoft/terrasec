package policy

const (
	azure                  supportedCloudType  = "azure"
	defaultAzureIacType    supportedIacType    = "terraform"
	defaultAzureIacVersion supportedIacVersion = version12
)

func init() {
	// Register azure as a cloud provider with terrasec
	RegisterCloudProvider(azure, defaultAzureIacType, defaultAzureIacVersion)
}
