

package mapper

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/core"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/cft"
)

// NewMapper returns a mapper based on IaC provider.
func NewMapper(iacType string) core.Mapper {
	switch iacType {
	case "cft":
		return cft.Mapper()
	case "arm":
		return arm.Mapper()
	}
	return nil
}
