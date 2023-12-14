package policy

const (
	github                  supportedCloudType  = "github"
	defaultGitHubIacType    supportedIacType    = "terraform"
	defaultGitHubIacVersion supportedIacVersion = version12
)

func init() {
	// Register github as a cloud provider with terrasec
	RegisterCloudProvider(github, defaultGitHubIacType, defaultGitHubIacVersion)
}
