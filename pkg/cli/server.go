

package cli

import (
	"github.com/spf13/cobra"
	httpserver "github.com/khulnasoft/terrasec/pkg/http-server"
)

var (
	// Port at which API server will listen
	port string

	// CertFile Certificate file path, required in order to enable secure HTTP server
	certFile string

	// PrivateKeyFile Private key file path, required in order to enable secure HTTP server
	privateKeyFile string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Terrasec as an API server",
	Long: `Terrasec

Run Terrasec as an API server that inspects incoming IaC (Infrastructure-as-Code) files and returns the scan results.
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		initial(cmd, args, true)
	},
	Run: server,
}

func server(cmd *cobra.Command, args []string) {
	httpserver.Start(port, certFile, privateKeyFile)
}

func init() {
	serverCmd.Flags().StringVarP(&privateKeyFile, "key-path", "", "", "private key file path")
	serverCmd.Flags().StringVarP(&certFile, "cert-path", "", "", "certificate file path")
	serverCmd.Flags().StringVarP(&port, "port", "p", httpserver.GatewayDefaultPort, "server port")
	RegisterCommand(rootCmd, serverCmd)
}
