package writer

import (
	"encoding/xml"
	"io"

	"go.uber.org/zap"
)

const (
	xmlFormat supportedFormat = "xml"
)

func init() {
	RegisterWriter(xmlFormat, XMLWriter)
}

// XMLWriter prints data in XML format
func XMLWriter(data interface{}, writers []io.Writer) error {
	j, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		zap.S().Errorf("failed to write XML output. error: '%v'", err)
		return err
	}
	for _, writer := range writers {
		writer.Write(j)
		writer.Write([]byte{'\n'})
	}
	return nil
}
