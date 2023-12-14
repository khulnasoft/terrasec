

package policy

const (
	aws                  supportedCloudType  = "aws"
	defaultAWSIacType    supportedIacType    = "terraform"
	defaultAWSIacVersion supportedIacVersion = version12
)

func init() {
	// Register aws as a cloud provider with terrasec
	RegisterCloudProvider(aws, defaultAWSIacType, defaultAWSIacVersion)
}
