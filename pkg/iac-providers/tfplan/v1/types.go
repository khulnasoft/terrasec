package tfplan

// TFPlan implements the IacProvider interface
type TFPlan struct {
	FormatVersion    string `json:"format_version"`
	TerraformVersion string `json:"terraform_version"`
}
