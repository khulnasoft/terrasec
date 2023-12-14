

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/rds"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// DBEventSubscriptionConfig holds config for aws_db_event_subscription resource
type DBEventSubscriptionConfig struct {
	Config
	SnsTopicArn     string   `json:"sns_topic"`
	Enabled         bool     `json:"enabled,omitempty"`
	EventCategories []string `json:"event_categories,omitempty"`
	SourceIds       []string `json:"source_ids,omitempty"`
	SourceType      string   `json:"source_type,omitempty"`
}

// GetDBEventSubscriptionConfig returns config for aws_db_event_subscription resource
func GetDBEventSubscriptionConfig(d *rds.EventSubscription) []AWSResourceConfig {
	cf := DBEventSubscriptionConfig{
		Config:          Config{},
		SnsTopicArn:     d.SnsTopicArn,
		Enabled:         functions.GetVal(d.Enabled),
		EventCategories: d.EventCategories,
		SourceIds:       d.SourceIds,
		SourceType:      functions.GetVal(d.SourceType),
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: d.AWSCloudFormationMetadata,
	}}
}
