package commons

import (
	"encoding/json"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
	"go.uber.org/zap"
)

var (
	// reference patterns
	varRefPattern = regexp.MustCompile(`(\$\{)?var\.(?P<name>\w*)(\})?`)
)

// isVarRef returns true if the given string is a variable reference
func isVarRef(attrVal string) bool {
	return varRefPattern.MatchString(attrVal)
}

// getVarName returns the actual variable name as configured in IaC. It trims
// of "${var." prefix and "}" suffix and returns the variable name
func getVarName(varRef string) (string, string) {

	// 1. extract the exact variable reference from the string
	varExpr := varRefPattern.FindString(varRef)

	// 2. extract variable name from variable reference
	match := varRefPattern.FindStringSubmatch(varRef)
	result := make(map[string]string)
	for i, name := range varRefPattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	varName := result["name"]

	zap.S().Debugf("extracted variable name %q from reference %q", varName, varRef)
	return varName, varExpr
}

// ResolveVarRef returns the variable value as configured in IaC config in module
func (r *RefResolver) ResolveVarRef(varRef, callerRef string) interface{} {

	// get variable name from varRef
	varName, varExpr := getVarName(varRef)

	// check if variable name exists in the map of variables read from IaC
	hclVar, present := r.Config.Module.Variables[varName]
	if !present {
		zap.S().Debugf("variable name: %q, ref: %q not present in variables", varName, varRef)
		return varRef
	}

	// return varRef if default value is not present, or value is a NilVal,
	// or if default value is not known
	if hclVar.Default.IsNull() || hclVar.Default.RawEquals(cty.NilVal) || !hclVar.Default.IsKnown() {
		return varRef
	}

	// default value is of cty.Value type, convert it to native golang type
	// based on cty.Type, determine golang type
	val, err := convertCtyToGoNative(hclVar.Default)
	if err != nil {
		zap.S().Debugf(err.Error())
		zap.S().Debugf("failed to convert cty.Value '%v' to golang native value", hclVar.Default.GoString())
		return varRef
	}
	zap.S().Debugf("resolved variable ref '%v', value: '%v'", varRef, val)

	valKind := reflect.TypeOf(val).Kind()

	if valKind == reflect.String || valKind == reflect.Map {
		valStr := ""

		if valKind == reflect.Map {
			data, err := json.Marshal(val)
			if err != nil {
				zap.S().Errorf("failed to convert expression '%v', ref: '%v'", hclVar, varRef)
				return varRef
			}
			valStr = string(data)
		} else {
			valStr = val.(string)
		}

		resolvedVal := strings.Replace(varRef, varExpr, valStr, 1)
		if varRef == resolvedVal {
			zap.S().Debugf("resolved str variable ref refers to self: '%v'", varRef)
			return varRef
		}
		if callerRef != "" && strings.Contains(resolvedVal, callerRef) {
			zap.S().Debugf("resolved str variable ref: '%v', value: '%v'", varRef, string(resolvedVal))
			return resolvedVal
		}
		return r.ResolveStrRef(resolvedVal, varRef)
	}
	return val
}

// ResolveVarRefFromParentModuleCall returns the variable value as configured in
// ModuleCall from parent module. The resolved value can be an absolute value
// (string, int, bool etc.) or it can also be another reference, which may
// need further resolution
func (r *RefResolver) ResolveVarRefFromParentModuleCall(varRef, callerRef string) interface{} {

	zap.S().Debugf("resolving variable ref %q in parent module call", varRef)

	// if module call struct is nil, nothing to process
	if r.ParentModuleCall == nil {
		return varRef
	}

	// get variable name from varRef
	varName, varExpr := getVarName(varRef)

	// get initialized variables from module call
	ParentModuleCallBody, ok := r.ParentModuleCall.Config.(*hclsyntax.Body)
	if !ok {
		return varRef
	}

	// get varName from module call, if present
	varAttr, present := ParentModuleCallBody.Attributes[varName]
	if !present {
		zap.S().Debugf("variable name: %q, ref: %q not present in parent module call", varName, varRef)
		return varRef
	}

	// read source file
	fileBytes, err := os.ReadFile(r.ParentModuleCall.SourceAddrRange.Filename)
	if err != nil {
		zap.S().Errorf("failed to read terraform IaC file '%s'. error: '%v'", r.ParentModuleCall.SourceAddr, err)
		return varRef
	}

	// extract values from attribute expressions as golang interface{}
	c := converter{bytes: fileBytes}
	val, _, err := c.convertExpression(varAttr.Expr)
	if err != nil {
		zap.S().Errorf("failed to convert expression '%v', ref: '%v'", varAttr.Expr, varRef)
		return varRef
	}

	// replace the variable reference string with actual value
	if reflect.TypeOf(val).Kind() == reflect.String {
		valStr := val.(string)

		// if resolved variable value from parent module is local reference get the local value from parent module
		if isLocalRef(valStr) {
			t := NewRefResolver(r.Config.Parent, nil)
			localVal := t.ResolveLocalRef(valStr, varRef)
			if reflect.TypeOf(localVal).Kind() == reflect.String {
				valStr = localVal.(string)
			}
		}

		resolvedVal := strings.Replace(varRef, varExpr, valStr, 1)
		if strings.Contains(valStr, varExpr) {
			zap.S().Debugf("resolved str variable ref refers to self: '%v'", varRef)
			return resolvedVal
		}
		if callerRef != "" && strings.Contains(resolvedVal, callerRef) {
			zap.S().Debugf("resolved str variable ref: '%v', value: '%v'", varRef, string(resolvedVal))
			return resolvedVal
		}
		return r.ResolveStrRef(resolvedVal, varRef)
	}

	// return extracted value
	zap.S().Debugf("resolved variable ref: '%v', value: '%v'", varRef, val)
	return val
}
