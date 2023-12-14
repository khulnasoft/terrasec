package config

import (
	"encoding/json"

	"github.com/awslabs/goformation/v7/cloudformation/s3"
)

// S3BucketPolicyConfig holds config for aws_s3_bucket_policy
type S3BucketPolicyConfig struct {
	Config
	PolicyDocument string `json:"policy"`
	Bucket         string `json:"bucket"`
}

// GetS3BucketPolicyConfig returns config for aws_s3_bucket_policy
func GetS3BucketPolicyConfig(p *s3.BucketPolicy) []AWSResourceConfig {
	cf := S3BucketPolicyConfig{
		Config: Config{
			Name: p.Bucket,
		},
		Bucket: p.Bucket,
	}

	policyDocument, err := json.Marshal(p.PolicyDocument)
	if err == nil {
		cf.PolicyDocument = string(policyDocument)
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: p.AWSCloudFormationMetadata,
	}}
}
