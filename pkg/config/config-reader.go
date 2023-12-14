

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	tomlExtension  = ".toml"
	yamlExtension1 = ".yaml"
	yamlExtension2 = ".yml"
)

var (
	// ErrTomlLoadConfig indicates error: Failed to load toml config
	ErrTomlLoadConfig = fmt.Errorf("failed to load toml config")
	// ErrNotPresent indicates error: Config file not present
	ErrNotPresent = fmt.Errorf("config file not present")
)

// TerrasecConfigReader holds the terrasec config file name
type TerrasecConfigReader struct {
	config TerrasecConfig
}

// NewTerrasecConfigReader initialises and returns a config reader
func NewTerrasecConfigReader(fileName string) (*TerrasecConfigReader, error) {
	config := TerrasecConfig{}
	configReader := new(TerrasecConfigReader)
	configReader.config = config

	// empty file name check should be done by the caller, this is a safe check
	if fileName == "" {
		zap.S().Debug("no config file specified")
		return configReader, nil
	}

	// return error if file doesn't exist
	_, err := os.Stat(fileName)
	if err != nil {
		zap.S().Errorf("config file: %s, doesn't exist", fileName)
		return configReader, ErrNotPresent
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		zap.S().Error("error loading config file", zap.Error(err))
		return configReader, ErrTomlLoadConfig
	}

	// check the extension of the file and decode using the file contents
	// using the relevant package
	switch filepath.Ext(fileName) {
	case tomlExtension:
		err = toml.Unmarshal(data, &configReader.config)
	case yamlExtension1, yamlExtension2:
		err = yaml.Unmarshal(data, &configReader.config)
	default:
		err = fmt.Errorf("file format %q not support for terrasec config file",
			filepath.Ext(fileName))
	}
	if err != nil {
		return configReader, err
	}

	return configReader, nil
}

// GetPolicyConfig will return the policy config from the terrasec config file
func (r TerrasecConfigReader) getPolicyConfig() Policy {
	return r.config.Policy
}

// GetNotifications will return the notifiers specified in the terrasec config file
func (r TerrasecConfigReader) getNotifications() map[string]Notifier {
	return r.config.Notifications
}

// GetRules will return the rules specified in the terrasec config file
func (r TerrasecConfigReader) getRules() Rules {
	return r.config.Rules
}

// GetCategory will return the category specified in the terrasec config file
func (r TerrasecConfigReader) getCategory() Category {
	return r.config.Category
}

// GetSeverity will return the level of severity specified in the terrasec config file
func (r TerrasecConfigReader) getSeverity() Severity {
	return r.config.Severity
}

// GetK8sAdmissionControl will return the k8s deny rules specified in the terrasec config file
func (r TerrasecConfigReader) GetK8sAdmissionControl() K8sAdmissionControl {
	return r.config.K8sAdmissionControl
}
