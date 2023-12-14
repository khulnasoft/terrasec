package functions

import (
	"strings"
)

// ResourceIDs is a map[ARMResource.Type]output.ResourceConfig.ID
// required for resolving the resourceId function calls in ARM templates.
var ResourceIDs = map[string]string{}

// LookUp function looks for different keywords in str
// and accordingly selects a function to call.
// generic variant, for eg. use for bool
func LookUp(vars, params map[string]interface{}, key string) interface{} {
	switch {
	case strings.Contains(key, "concat("):
		return Concat(vars, params, key)
	case strings.Contains(key, "toLower("):
		return ToLower(vars, params, key)
	case strings.Contains(key, "resourceId("):
		return ResourceID(vars, params, key)
	case strings.Contains(key, "parameters("):
		s := Parameters(key)
		return LookUp(vars, params, s)
	case strings.Contains(key, "variables("):
		s := Variables(key)
		return LookUp(vars, params, s)
	case strings.Contains(key, "uniqueString("):
		return UniqueString()
	default:
		if v, ok := params[key]; ok {
			if res, ok := v.(string); ok {
				if strings.Contains(res, key) {
					return key
				}
				return LookUp(vars, params, res)
			}
			return v
		}
		if v, ok := vars[key]; ok {
			if res, ok := v.(string); ok {
				if strings.Contains(res, key) {
					return key
				}
				return LookUp(vars, params, res)
			}
			return v
		}
		return key
	}
}

// LookUpFloat64 safely returns float64 after Lookup
func LookUpFloat64(vars, params map[string]interface{}, key string) float64 {
	if value, ok := LookUp(vars, params, key).(float64); ok {
		return value
	}
	return 0
}

// LookUpString safely returns string after Lookup
func LookUpString(vars, params map[string]interface{}, key string) string {
	if value, ok := LookUp(vars, params, key).(string); ok {
		return value
	}
	return key
}
