package config

import (
	"github.com/khulnasoft/terrasec/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrasec/pkg/mapper/iac-providers/arm/types"
)

const (
	armDayOfWeek       = "dayOfWeek"
	armStartHourUtc    = "startHourUtc"
	armScheduleEntries = "scheduleEntries"
)

const (
	tfDayOfWeek    = "day_of_week"
	tfStartHourUTC = "start_hour_utc"
)

// PatchScheduleConfig returns config for patch_schedule
func PatchScheduleConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	sch := convert.ToMap(r.Properties, armScheduleEntries)
	return map[string]interface{}{
		tfDayOfWeek:    fn.LookUpString(nil, params, convert.ToString(sch, armDayOfWeek)),
		tfStartHourUTC: fn.LookUpFloat64(nil, params, convert.ToString(sch, armStartHourUtc)),
		tfTags:         r.Tags,
	}
}
