

package writer

import (
	"encoding/json"
	"io"
)

const (
	jsonFormat supportedFormat = "json"
)

func init() {
	RegisterWriter(jsonFormat, JSONWriter)
}

// JSONWriter prints data in JSON format
func JSONWriter(data interface{}, writers []io.Writer) error {
	j, _ := json.MarshalIndent(data, "", "  ")
	for _, writer := range writers {
		writer.Write(j)
		writer.Write([]byte{'\n'})
	}
	return nil
}
