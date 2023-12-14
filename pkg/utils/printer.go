

package utils

import (
	"encoding/json"
	"io"
)

// PrintJSON prints data in JSON format
func PrintJSON(data interface{}, writer io.Writer) {
	j, _ := json.MarshalIndent(data, "", "  ")
	writer.Write(j)
	writer.Write([]byte{'\n'})
}
