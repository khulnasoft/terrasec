

package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	iacProvider "github.com/khulnasoft/terrasec/pkg/iac-providers"
	"go.uber.org/zap"
)

// IacProvider contains response body for iac providers
type IacProvider struct {
	Type           string   `json:"type"`
	Versions       []string `json:"versions"`
	DefaultVersion string   `json:"defaultVersion"`
}

// iacProviders returns list of iac providers
func (g *APIHandler) iacProviders(w http.ResponseWriter, r *http.Request) {
	var providers = []IacProvider{}
	for _, provider := range iacProvider.SupportedIacProviders() {
		providers = append(providers, IacProvider{
			Type:           string(provider),
			Versions:       iacProvider.GetProviderIacVersions(provider),
			DefaultVersion: iacProvider.GetDefaultIacVersion(provider),
		})
	}

	response, err := json.MarshalIndent(providers, "", "  ")
	if err != nil {
		errMsg := fmt.Sprintf("failed to create JSON. error: '%v'", err)
		zap.S().Error(errMsg)
		apiErrorResponse(w, errMsg, http.StatusInternalServerError)
		return
	}

	apiResponse(w, string(response), http.StatusOK)
}
