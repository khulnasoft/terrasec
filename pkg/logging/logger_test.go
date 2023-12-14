

package logging

import (
	"reflect"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestGetLoggerLevel(t *testing.T) {
	table := []struct {
		name  string
		input string
		want  zapcore.Level
	}{
		{
			name:  "empty log level",
			input: "",
			want:  zapcore.InfoLevel,
		},
		{
			name:  "invalid log level",
			input: "some log level",
			want:  zapcore.InfoLevel,
		},
		{
			name:  "debug log level",
			input: "debug",
			want:  zapcore.DebugLevel,
		},
		{
			name:  "panic log level",
			input: "panic",
			want:  zapcore.PanicLevel,
		},
	}

	for _, tt := range table {
		got := getLoggerLevel(tt.input)
		if got != tt.want {
			t.Errorf("got: '%v', want: '%v'", got, tt.want)
		}
	}
}

func TestGetLogger(t *testing.T) {

	table := []struct {
		name          string
		logLevel      string
		encoding      string
		encodingLevel func(zapcore.Level, zapcore.PrimitiveArrayEncoder)
		logDir        string
	}{
		{
			name:          "check debug log level",
			logLevel:      "debug",
			encoding:      "json",
			encodingLevel: zapcore.LowercaseLevelEncoder,
			logDir:        "",
		},
		{
			name:          "check log level",
			logLevel:      "panic",
			encoding:      "console",
			encodingLevel: zapcore.LowercaseLevelEncoder,
			logDir:        "",
		},
	}

	for _, tt := range table {
		got := GetLogger(tt.logLevel, tt.encoding, tt.logDir, tt.encodingLevel)
		if ce := got.Check(getLoggerLevel(tt.logLevel), "testing"); ce == nil {
			t.Errorf("unexpected error")
		}
	}
}

func TestGetDefaultLogger(t *testing.T) {
	t.Run("json encoding", func(t *testing.T) {
		Init("json", "info", "")
		got := GetDefaultLogger()
		want := globalLogger
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})
	t.Run("console encoding", func(t *testing.T) {
		Init("console", "info", "")
		got := GetDefaultLogger()
		want := globalLogger
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})
}
