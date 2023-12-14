

package cli

import (
	"github.com/spf13/cobra"
	"github.com/khulnasoft/terrasec/pkg/initialize"
	"go.uber.org/zap"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes Terrasec and clones policies from the Terrasec GitHub repository.",
	Long: `Terrasec

Initializes Terrasec and clones policies from the Terrasec GitHub repository.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return initial(cmd, args, false)
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func initial(cmd *cobra.Command, args []string, nonInitCmd bool) error {
	// initialize terrasec
	if err := initialize.Run(nonInitCmd); err != nil {
		zap.S().Errorf("failed to initialize terrasec. error : %v", err)
		return err
	}
	return nil
}

func init() {
	RegisterCommand(rootCmd, initCmd)
}
