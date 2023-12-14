package runtime

import (
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/policy"
)

// Output is the runtime engine output
type Output struct {
	ResourceConfig output.AllResourceConfigs
	Violations     policy.EngineOutput
}

// engineEvalResult represents engine evaluation result
type engineEvalResult struct {
	err    error
	output policy.EngineOutput
}

// dirScanResp represents a directory scan response
type dirScanResp struct {
	err     error
	rc      output.AllResourceConfigs
	iacType string
}
