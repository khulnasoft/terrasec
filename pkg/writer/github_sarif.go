package writer

import (
	"io"
)

const (
	githubSarifFormat supportedFormat = "github-sarif"
)

func init() {
	RegisterWriter(githubSarifFormat, GitHubSarifWriter)
}

// GitHubSarifWriter writes sarif formatted violation results report that are well suited for github codescanning alerts display
func GitHubSarifWriter(data interface{}, writers []io.Writer) error {
	return writeSarif(data, writers, true)
}
