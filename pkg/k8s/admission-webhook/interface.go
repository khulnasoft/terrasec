package admissionwebhook

import (
	admissionv1 "k8s.io/api/admission/v1"
)

// AdmissionWebhook interface needs to be implemented by all k8s admission
// webhooks i.e validating and mutating webhooks
type AdmissionWebhook interface {

	// Authorize checks if the incoming webhooks have valid apiKey
	Authorize(apiKey string) error

	// DecodeAdmissionReviewRequest reads the incoming admission request body
	// and decodes it into an AdmissionReviewRequest struct
	DecodeAdmissionReviewRequest(payload []byte) (admissionv1.AdmissionReview, error)

	// ProcessWebhook processes the incoming AdmissionReview and creates
	// a AdmissionResponse
	ProcessWebhook(review admissionv1.AdmissionReview, serverURL string) (*admissionv1.AdmissionReview, error)
}
