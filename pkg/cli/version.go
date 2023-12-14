package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/terrasec/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Terrasec version",
	Long: `Terrasec

Displays the version of this Terrasec binary
`,
	Run: getVersion,
}

func getVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("version: %s\n", version.Get())
}

func init() {
	RegisterCommand(rootCmd, versionCmd)
}