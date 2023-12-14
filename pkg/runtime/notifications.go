package runtime

import (
	"github.com/khulnasoft/terrasec/pkg/utils"
)

// SendNotifications sends notifications via all the configured notifiers
func (e *Executor) SendNotifications(data interface{}) error {
	var allErrs error

	// send notifications using configured notifiers
	for _, notifier := range e.notifiers {
		err := notifier.SendNotification(data)
		if err != nil {
			allErrs = utils.WrapError(err, allErrs)
			continue
		}
	}
	return allErrs
}
