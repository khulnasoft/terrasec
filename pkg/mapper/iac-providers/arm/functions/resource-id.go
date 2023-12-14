

package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// ResourceID function runs str against a regular
// expression and returns the resource ID.
//
// For example:
// if str = [resourceId('Microsoft.KeyVault/vaults', parameters('keyVaultName'))],
// the function returns resource ID for Microsoft.KeyVault/vaults.
func ResourceID(vars, params map[string]interface{}, str string) string {
	const (
		start = "resourceId("
		end   = ")"
	)

	key := strings.TrimPrefix(str, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)

	if len(results) == 0 {
		zap.S().Debugf("failed to parse expression: %s", str)
		return ""
	}

	rs := strings.Split(results[0][1], ",")
	if id, ok := ResourceIDs[strings.Trim(rs[0], "'")]; ok {
		return id
	}
	return ""
}
