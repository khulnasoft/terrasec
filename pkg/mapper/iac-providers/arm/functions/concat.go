package functions

import (
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
	"go.uber.org/zap"
)

// Concat function splits str and runs respective functions on split parts.
// Example: [Concat(parameters('vaultName'), '/', parameters('keyName'))]
func Concat(vars, params map[string]interface{}, str string) string {
	const (
		start = "concat("
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

	sb := &strings.Builder{}
	cs := strings.Split(results[0][1], ",")
	for _, s := range cs {
		s = strings.TrimSpace(s)
		s = strings.Trim(s, "'")
		if _, err := sb.WriteString(LookUpString(vars, params, s)); err != nil {
			zap.S().Debugf("failed to parse expression: %s", str)
			return ""
		}
	}
	return sb.String()
}
