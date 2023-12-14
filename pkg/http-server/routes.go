package httpserver

import (
	"net/http"
)

// Route is a specification and  handler for a REST endpoint.
type Route struct {
	verb string
	path string
	fn   func(http.ResponseWriter, *http.Request)
}

// Routes returns a slice of routes of API endpoints to be registered with
// http server
func (g *APIServer) Routes() []*Route {
	h := NewAPIHandler()
	routes := []*Route{
		{verb: "GET", path: "/health", fn: h.Health},
		{verb: "GET", path: versionedPath("/providers"), fn: h.iacProviders},
		{verb: "POST", path: versionedPath("/{iac}/{iacVersion}/{cloud}/local/file/scan"), fn: h.scanFile},
		{verb: "POST", path: versionedPath("/{iac}/{iacVersion}/{cloud}/remote/dir/scan"), fn: h.scanRemoteRepo},

		// k8s webhook Routes
		{verb: "POST", path: versionedPath("/k8s/webhooks/{apiKey}/scan/validate"), fn: h.validateK8SWebhook},
	}

	return routes
}

func versionedPath(route string) string {
	return "/" + APIVersion + route
}
