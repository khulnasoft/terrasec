package filters

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/policy"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
)

// PolicyTypesFilterSpecification is policy type based Filter Spec
type PolicyTypesFilterSpecification struct {
	policyTypes []string
}

// IsSatisfied implementation for policy type based Filter spec
func (p PolicyTypesFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	// if policy type is not present for rego metadata,
	// or if policy types is not specified, return true
	if len(r.PolicyType) < 1 || len(p.policyTypes) < 1 {
		return true
	}
	return utils.CheckPolicyType(r.PolicyType, p.policyTypes)
}

// ResourceTypeFilterSpecification is resource type based Filter Spec
type ResourceTypeFilterSpecification struct {
	resourceType string
}

// IsSatisfied implementation for resource type based Filter spec
func (rs ResourceTypeFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	// if resource type is not present for rego metadata, return true
	if len(r.ResourceType) < 1 {
		return true
	}

	// temporary fix for kubernetes policies (https://github.com/khulnasoft/terrasec/issues/946)
	// TODO: modification in policy metadata (resource type (string) => resource type (map<string,bool))
	// accordingly modify the filter specification
	if r.PolicyType == "k8s" && strings.Contains(strings.ToLower(rs.resourceType), "kubernetes") {
		return true
	}

	return rs.resourceType == r.ResourceType
}

// ReferenceIDFilterSpecification is reference ID based Filter Spec
type ReferenceIDFilterSpecification struct {
	ReferenceID string
}

// IsSatisfied implementation for reference ID based Filter spec
func (rs ReferenceIDFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	if rs.ReferenceID == r.ID {
		return true
	}
	if rs.ReferenceID == r.ReferenceID {
		zap.S().Warnf("Deprecation warning : Use 'id' (%s) instead of 'reference_id' (%s) to skip/scan rules", r.ID, r.ReferenceID)
		return true
	}
	return false
}

// ReferenceIDsFilterSpecification is reference IDs based Filter Spec
type ReferenceIDsFilterSpecification struct {
	ReferenceIDs []string
}

// IsSatisfied implementation for reference IDs based Filter spec
func (rs ReferenceIDsFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	// when reference ID's are not specified (could be skip or scan rules),
	// return true
	if len(rs.ReferenceIDs) < 1 {
		return true
	}
	isSatisfied := false
	for _, refID := range rs.ReferenceIDs {
		rfIDSpec := ReferenceIDFilterSpecification{refID}
		if rfIDSpec.IsSatisfied(r) {
			isSatisfied = true
			break
		}
	}
	return isSatisfied
}

// CategoryFilterSpecification is categories based Filter Spec
type CategoryFilterSpecification struct {
	categories []string
}

// IsSatisfied implementation for category based Filter spec
func (c CategoryFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	// when categories are not specified, return true
	if len(c.categories) < 1 {
		return true
	}
	return utils.CheckCategory(r.Category, c.categories)
}

// SeverityFilterSpecification is severity based Filter Spec
type SeverityFilterSpecification struct {
	severity string
}

// IsSatisfied implementation for severity based Filter spec
func (s SeverityFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	// when severity is not specified, return true
	if len(s.severity) < 1 {
		return true
	}
	return utils.CheckSeverity(r.Severity, s.severity)
}

// AndFilterSpecification is a logical AND Filter spec which
// determines if a list of filter specs satisfy the condition
type AndFilterSpecification struct {
	filterSpecs []policy.FilterSpecification
}

// IsSatisfied implementation for And Filter spec
func (a AndFilterSpecification) IsSatisfied(r *policy.RegoMetadata) bool {
	if len(a.filterSpecs) < 1 {
		return false
	}
	for _, filterSpec := range a.filterSpecs {
		if !filterSpec.IsSatisfied(r) {
			return false
		}
	}
	return true
}
