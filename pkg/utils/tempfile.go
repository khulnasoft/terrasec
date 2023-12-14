

package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// CreateTempFile creates a file with provided contents in the temp directory
func CreateTempFile(content []byte, ext string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", fmt.Sprintf("terrasec-*.%s", ext))
	if err != nil {
		zap.S().Errorf("failed to create temp file: '%v'", err)
		return nil, err
	}

	zap.S().Debugf("created temp config file at '%s'", tempFile.Name())

	_, err = tempFile.Write(content)
	if err != nil {
		zap.S().Errorf("failed to write to temp file: '%v'", err)
		return nil, err
	}

	return tempFile, nil
}
