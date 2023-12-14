package validatingwebhook

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/khulnasoft/terrasec/pkg/config"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"github.com/pelletier/go-toml"
)

// CreateTerrasecConfigFile creates a config file with test policy path
func CreateTerrasecConfigFile(configFileName, policyRootRelPath string, terrasecConfig *config.TerrasecConfig) error {
	policyAbsPath, err := filepath.Abs(policyRootRelPath)
	if err != nil {
		return err
	}

	if utils.IsWindowsPlatform() {
		policyAbsPath = strings.ReplaceAll(policyAbsPath, "\\", "\\\\")
	}

	if terrasecConfig == nil {
		terrasecConfig = &config.TerrasecConfig{}
	}

	terrasecConfig.BasePath = policyAbsPath
	terrasecConfig.RepoPath = policyAbsPath

	// create config file in work directory
	file, err := os.Create(configFileName)
	if err != nil {
		return fmt.Errorf("config file creation failed, err: %v", err)
	}

	contentBytes, err := toml.Marshal(terrasecConfig)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(contentBytes))
	if err != nil {
		return fmt.Errorf("error while writing to config file, err: %v", err)
	}
	return nil
}

// CreateCertificate creates certificates required to run server in the folder specified
func CreateCertificate(certsFolder, certFileName, privKeyFileName string) (string, string, error) {
	// create certs folder to keep certificates
	os.Mkdir(certsFolder, 0755)
	certFileAbsPath, err := filepath.Abs(filepath.Join(certsFolder, "server.crt"))
	if err != nil {
		return "", "", err
	}
	privKeyFileAbsPath, err := filepath.Abs(filepath.Join(certsFolder, "priv.key"))
	if err != nil {
		return "", "", err
	}
	err = GenerateCertificates(certFileAbsPath, privKeyFileAbsPath)
	if err != nil {
		return "", "", err
	}

	return certFileAbsPath, privKeyFileAbsPath, nil
}

// DeleteDefaultKindCluster deletes the default kind cluster
func DeleteDefaultKindCluster() error {
	cmd := exec.Command("kind", "delete", "cluster")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// CreateDefaultKindCluster creates the default kind cluster
func CreateDefaultKindCluster() error {
	cmd := exec.Command("kind", "create", "cluster")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// GetIP finds preferred outbound ip of the machine
func GetIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4(), nil
			}
		}
	}
	return nil, fmt.Errorf("could not find ip address of the machine")
}
