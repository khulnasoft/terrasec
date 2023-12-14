package config

import (
	"strings"

	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armEmails              = "emails"
	armPhone               = "phone"
	armAlertNotifications  = "alertNotifications"
	armNotificationsByRole = "notificationsByRole"
	armState               = "state"
)

const (
	tfEmail              = "email"
	tfPhone              = "phone"
	tfAlertNotifications = "alert_notifications"
	tfAlertsToAdmins     = "alerts_to_admins"
)

// SecurityCenterContactConfig returns config for azurerm_security_center_contact
func SecurityCenterContactConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
		tfPhone:    fn.LookUpString(nil, params, convert.ToString(r.Properties, armPhone)),
		tfEmail:    fn.LookUpString(nil, params, convert.ToString(r.Properties, armEmails)),
	}

	notifications := convert.ToMap(r.Properties, armAlertNotifications)
	state := convert.ToString(notifications, armState)
	cf[tfAlertNotifications] = strings.EqualFold(strings.ToUpper(state), "ON")

	notifications = convert.ToMap(r.Properties, armNotificationsByRole)
	state = convert.ToString(notifications, state)
	cf[tfAlertsToAdmins] = strings.EqualFold(strings.ToUpper(state), "ON")

	return cf
}
