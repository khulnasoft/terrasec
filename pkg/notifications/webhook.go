

package notifications

import (
	"reflect"

	webhookNotifier "github.com/khulnasoft/terrasec/pkg/notifications/webhook"
)

// terraform specific constants
const (
	terraform supportedNotifierType = "webhook"
)

// register terraform as an IaC provider with terrasec
func init() {

	// register iac provider
	RegisterNotifier(terraform, reflect.TypeOf(webhookNotifier.Webhook{}))
}
