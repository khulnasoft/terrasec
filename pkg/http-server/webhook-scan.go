package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	admissionWebhook "github.com/khulnasoft/terrasec/pkg/k8s/admission-webhook"
	"go.uber.org/zap"

	v1 "k8s.io/api/admission/v1"
)

// validateK8SWebhook handles the incoming validating admission webhook from kubernetes API server
func (g *APIHandler) validateK8SWebhook(w http.ResponseWriter, r *http.Request) {
	zap.S().Debug("handle: validating webhook request")

	var (
		params                   = mux.Vars(r)
		apiKey                   = params["apiKey"]
		qP                       = r.URL.Query()
		notificationWebhookURL   = qP.Get("webhook-url")
		notificationWebhookToken = qP.Get("webhook-token")
		repoURL                  = qP.Get("repo-url")
		repoRef                  = qP.Get("repo-ref")
	)

	// Read the request into byte array
	body, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("failed to read validating admission webhook request body, error: '%v'", err)
		apiErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	zap.S().Debugf("scanning configuration webhook request: %+v", string(body))

	validatingWebhook := admissionWebhook.NewValidatingWebhook(body, notificationWebhookURL, notificationWebhookToken, repoURL, repoRef)
	// Validate if authorized (API key is specified and matched the server one (saved in an environment variable)
	if err := validatingWebhook.Authorize(apiKey); err != nil {
		switch err {
		case admissionWebhook.ErrAPIKeyMissing:
			apiErrorResponse(w, err.Error(), http.StatusBadRequest)
		case admissionWebhook.ErrUnauthorized:
			apiErrorResponse(w, err.Error(), http.StatusUnauthorized)
		default:
			apiErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// decode incoming admission review request
	requestedAdmissionReview, err := validatingWebhook.DecodeAdmissionReviewRequest(body)
	if err != nil {
		apiErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestedAdmissionReview.Request == nil {
		apiErrorResponse(w, "empty validating admission review request", http.StatusBadRequest)
		return
	}

	// process the admission review request
	admissionResponse, err := validatingWebhook.ProcessWebhook(requestedAdmissionReview, r.Host)
	if err != nil && err != admissionWebhook.ErrEmptyAdmissionReview {
		apiErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the correct response according to the result
	g.sendResponseAdmissionReview(w, admissionResponse)
}

func (g *APIHandler) sendResponseAdmissionReview(w http.ResponseWriter, admissionResponse *v1.AdmissionReview) {
	respBytes, err := json.Marshal(admissionResponse)
	if err != nil {
		msg := fmt.Sprintf("failed to serialize admission review response: %v", err)
		zap.S().Error(msg)
		apiErrorResponse(w, msg, http.StatusInternalServerError)
		return
	}

	zap.S().Debugf("response result: %+v", string(respBytes))
	apiResponse(w, string(respBytes), http.StatusOK)
}
