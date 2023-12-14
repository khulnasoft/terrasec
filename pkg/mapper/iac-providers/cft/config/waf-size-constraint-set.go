

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/waf"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft/functions"
)

// FieldToMatchBlock holds field_to_match attribute
type FieldToMatchBlock struct {
	Data string `json:"data,omitempty"`
	Type string `json:"type"`
}

// SizeConstraintSetBlock holds size_constraints attribute
type SizeConstraintSetBlock struct {
	ComparisonOperator string              `json:"comparison_operator"`
	Size               int                 `json:"size"`
	TextTransformation string              `json:"text_transformation"`
	FieldToMatch       []FieldToMatchBlock `json:"field_to_match"`
}

// WafSizeConstraintSetConfig holds Config for aws_waf_size_constraint_set
type WafSizeConstraintSetConfig struct {
	Config
	Name              string                   `json:"name"`
	SizeConstraintSet []SizeConstraintSetBlock `json:"size_constraints,omitempty"`
}

// GetWafSizeConstraintSetConfig returns config for aws_waf_size_constraint_set
func GetWafSizeConstraintSetConfig(w *waf.SizeConstraintSet) []AWSResourceConfig {
	sizeConstraintSet := setSizeConstraintSet(w)

	cf := WafSizeConstraintSetConfig{
		Config: Config{
			Name: w.Name,
		},
		Name:              w.Name,
		SizeConstraintSet: sizeConstraintSet,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: w.AWSCloudFormationMetadata,
	}}
}

func setSizeConstraintSet(w *waf.SizeConstraintSet) []SizeConstraintSetBlock {
	sizeConstraintSet := make([]SizeConstraintSetBlock, len(w.SizeConstraints))

	for i := range w.SizeConstraints {
		sizeConstraintSet[i].Size = w.SizeConstraints[i].Size
		sizeConstraintSet[i].TextTransformation = w.SizeConstraints[i].TextTransformation
		sizeConstraintSet[i].ComparisonOperator = w.SizeConstraints[i].ComparisonOperator
		sizeConstraintSet[i].FieldToMatch = setFieldToMatch(w, i)
	}

	return sizeConstraintSet
}

func setFieldToMatch(w *waf.SizeConstraintSet, index int) []FieldToMatchBlock {
	fieldToMatchBlock := make([]FieldToMatchBlock, 1)
	fieldToMatchBlock[0].Type = w.SizeConstraints[index].FieldToMatch.Type

	if w.SizeConstraints[index].FieldToMatch.Data != nil {
		fieldToMatchBlock[0].Data = functions.GetVal(w.SizeConstraints[index].FieldToMatch.Data)
	}

	return fieldToMatchBlock
}
