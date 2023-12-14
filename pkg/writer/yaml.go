

package writer

import (
	"io"

	"gopkg.in/yaml.v3"
)

const (
	yamlFormat supportedFormat = "yaml"
)

func init() {
	RegisterWriter(yamlFormat, YAMLWriter)
}

// YAMLWriter prints data in YAML format
func YAMLWriter(data interface{}, writers []io.Writer) error {
	j, _ := yaml.Marshal(data)
	for _, writer := range writers {
		writer.Write(j)
		writer.Write([]byte{'\n'})
	}
	return nil
}
