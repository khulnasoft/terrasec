package writer

import (
	"io"
)

// supportedFormat data type for supported formats
type supportedFormat string

// writerMap stores mapping of supported writer formats with respective functions
var writerMap = make(map[supportedFormat](func(interface{}, []io.Writer) error))

// RegisterWriter registers a writer for terrasec
func RegisterWriter(format supportedFormat, writerFunc func(interface{}, []io.Writer) error) {
	writerMap[format] = writerFunc
}
