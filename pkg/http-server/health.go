

package httpserver

import (
	"net/http"

	"go.uber.org/zap"
)

// Health returns the health of the http server
func (g *APIHandler) Health(w http.ResponseWriter, r *http.Request) {
	zap.S().Debug("handle: health check request")
	w.WriteHeader(http.StatusOK)
}
