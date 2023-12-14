package writer

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

var (
	errNotSupported = fmt.Errorf("output format not supported")
)

// Write method writes in the given format using the respective writer func
func Write(format string, data interface{}, writers []io.Writer) error {

	writerFunc, present := writerMap[supportedFormat(format)]
	if !present {
		zap.S().Errorf("output format '%s' not supported", format)
		return errNotSupported
	}

	return writerFunc(data, writers)
}
