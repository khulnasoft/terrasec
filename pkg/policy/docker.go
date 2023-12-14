package policy

const (
	docker                  supportedCloudType  = "docker"
	defaultDockerIacType    supportedIacType    = "docker"
	defaultDockerIacVersion supportedIacVersion = version1
)

func init() {
	// Register docker as a cloud provider with terrasec
	RegisterCloudProvider(docker, defaultDockerIacType, defaultDockerIacVersion)
}
