package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	validJSON      []byte
	validJSONInput = map[string]int{"apple": 5, "lettuce": 7}
	validJSONFile  = filepath.Join(testDataDir, "valid.json")
)

func init() {
	validJSON, _ = os.ReadFile(validJSONFile)

}

func TestPrintJSON(t *testing.T) {

	table := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name:  "empty JSON",
			input: make(map[string]interface{}),
			want:  "{}",
		},
		{
			name:  "valid JSON",
			input: validJSONInput,
			want:  string(validJSON),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got := &bytes.Buffer{}
			PrintJSON(tt.input, got)
			if IsWindowsPlatform() {
				tt.want = ReplaceWinNewLineString(tt.want)
			}
			if strings.TrimSpace(got.String()) != strings.TrimSpace(tt.want) {
				t.Errorf("got:\n'%v'\n, want:\n'%v'\n", got, tt.want)
			}
		})
	}
}
