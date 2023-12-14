package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// ToLower function runs str against a regular expression
// and returns the final value in all lower case.
//
// For example:
// if param = [toLower('location')],
// the function returns location as the key.
func ToLower(vars, params map[string]interface{}, str string) string {
	const (
		start = "toLower("
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
	return strings.ToLower(LookUp(vars, params, results[0][1]).(string))
}
