

package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// Variables function runs variable against a regular
// expression and return the variable key.
//
// For example:
// if var = Variables('identityName'),
// the function returns identityName as the key.
func Variables(variable string) string {
	const (
		start = "variables('"
		end   = "')"
	)

	key := strings.TrimPrefix(variable, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)

	if len(results) == 0 {
		zap.S().Debugf("failed to parse expression: %s", variable)
		return ""
	}
	return results[0][1]
}
