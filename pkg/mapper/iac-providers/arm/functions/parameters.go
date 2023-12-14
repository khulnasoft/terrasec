

package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// Parameters function runs param against a regular
// expression and returns the parameter key.
//
// For example:
// if param = [Parameters('location')],
// the function returns location as the key.
func Parameters(param string) string {
	const (
		start = "parameters('"
		end   = "')"
	)

	key := strings.TrimPrefix(param, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)
	if len(results) == 0 {
		zap.S().Debugf("failed to parse expression: %s", param)
		return ""
	}
	return results[0][1]
}
