package httpserver

import "fmt"

func (g *APIServer) validateFiles(privateKeyFile, certFile string) error {
	keylength := len(privateKeyFile)
	certlength := len(certFile)

	if keylength > 0 && certlength == 0 {
		return fmt.Errorf("private key file provided but certificate file missing")
	} else if keylength == 0 && certlength > 0 {
		return fmt.Errorf("certificate file provided but private key file missing")
	}

	return nil
}
