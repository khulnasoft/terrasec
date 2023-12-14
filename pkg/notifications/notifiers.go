

package notifications

import (
	"fmt"
	"reflect"

	"github.com/khulnasoft/terrasec/pkg/config"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
)

var (
	errNotifierNotSupported   = fmt.Errorf("notifier not supported")
	errNotifierTypeNotPresent = fmt.Errorf("notifier type not present in toml config")
	// ErrNotificationNotPresent error is caused when there isn't any notification present in the config
	ErrNotificationNotPresent = fmt.Errorf("no notification specified in the config")
)

// NewNotifier returns a new notifier
func NewNotifier(notifierType string) (notifier Notifier, err error) {

	// get notifier from supportedNotifiers
	notifierObject, supported := supportedNotifiers[supportedNotifierType(notifierType)]
	if !supported {
		zap.S().Errorf("notifier type '%s' not supported", notifierType)
		return notifier, errNotifierNotSupported
	}

	// successful
	return reflect.New(notifierObject).Interface().(Notifier), nil
}

// NewNotifiers returns a list of notifiers configured in the config file
func NewNotifiers() ([]Notifier, error) {

	var notifiers []Notifier

	// get config for 'notifications'
	notifications := config.GetNotifications()
	if len(notifications) == 0 {
		return notifiers, ErrNotificationNotPresent
	}

	// create notifiers
	var allErrs error
	for _, notifier := range notifications {
		if notifier.NotifierType == "" {
			zap.S().Error(errNotifierTypeNotPresent)
			allErrs = utils.WrapError(errNotifierTypeNotPresent, allErrs)
			continue
		}

		if !IsNotifierSupported(notifier.NotifierType) {
			zap.S().Errorf("notifier type '%s' not supported", notifier.NotifierType)
			allErrs = utils.WrapError(errNotifierNotSupported, allErrs)
			continue
		}

		// create a new notifier
		n, err := NewNotifier(notifier.NotifierType)
		if err != nil {
			allErrs = utils.WrapError(err, allErrs)
			continue
		}

		// populate data
		err = n.Init(notifier.NotifierConfig)
		if err != nil {
			allErrs = utils.WrapError(err, allErrs)
			continue
		}

		// add to the list of notifiers
		notifiers = append(notifiers, n)
	}

	// return list of notifiers
	return notifiers, allErrs
}

// IsNotifierSupported returns true/false depending on whether the notifier
// is supported in terrasec or not
func IsNotifierSupported(notifierType string) bool {
	if _, supported := supportedNotifiers[supportedNotifierType(notifierType)]; !supported {
		return false
	}
	return true
}
