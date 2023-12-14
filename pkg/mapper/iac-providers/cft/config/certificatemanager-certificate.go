

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/certificatemanager"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// CertificateManagerCertificateConfig holds config for CertificateManagerCertificate
type CertificateManagerCertificateConfig struct {
	Config
	DomainName       string `json:"domain_name"`
	ValidationMethod string `json:"validation_method"`
}

// GetCertificateManagerCertificateConfig returns config for CertificateManagerCertificate
func GetCertificateManagerCertificateConfig(c *certificatemanager.Certificate) []AWSResourceConfig {
	cf := CertificateManagerCertificateConfig{
		Config: Config{
			Tags: c.Tags,
		},
		DomainName:       c.DomainName,
		ValidationMethod: functions.GetVal(c.ValidationMethod),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: c.AWSCloudFormationMetadata,
	}}
}
