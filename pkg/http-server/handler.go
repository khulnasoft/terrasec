

package httpserver

// APIHandler struct for http api server
type APIHandler struct {
	test bool
}

// NewAPIHandler returns a new APIHandler{}
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}
