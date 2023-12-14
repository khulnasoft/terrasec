

package cli

import (
	"github.com/spf13/cobra"
)

var (
	// LogLevel Logging level (debug, info, warn, error, panic, fatal)
	LogLevel string

	// LogType Logging output type (console, json)
	LogType string

	// OutputType Violation output type (human, json, yaml, xml, sarif)
	OutputType string

	// ConfigFile Config file path
	ConfigFile string

	// CustomTempDir Temporary directory path to download remote repository,module and templates
	CustomTempDir string

	// LogOutputDir Directory to write scan logs and result files
	LogOutputDir string
)

var rootCmd = &cobra.Command{
	Use:   "terrasec",
	Short: "Detect compliance and security violations across Infrastructure as Code.",
	Long: `Terrasec

Detect compliance and security violations across Infrastructure as Code to mitigate risk before provisioning cloud native infrastructure.
For more information, please visit https://runterrasec.io/
`,
}
