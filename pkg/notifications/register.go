

package notifications

import (
	"reflect"
)

// map of supported notifier types
var supportedNotifiers = make(map[supportedNotifierType]reflect.Type)

// RegisterNotifier registers an notifier provider for terrasec
func RegisterNotifier(notifierType supportedNotifierType, notifierProvider reflect.Type) {
	supportedNotifiers[notifierType] = notifierProvider
}
