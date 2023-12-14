package httpserver

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateFiles(t *testing.T) {
	server := APIServer{}
	table := []struct {
		name           string
		privateKeyFile string
		certFile       string
		wantOutput     interface{}
		wantErr        error
	}{
		{
			name:           "both file names provided",
			privateKeyFile: "key",
			certFile:       "cert",
			wantErr:        nil,
		},
		{
			name:           "privatekey filename absent",
			privateKeyFile: "",
			certFile:       "server.crt",
			wantErr:        fmt.Errorf("certificate file provided but private key file missing"),
		},
		{
			name:           "both file names blank",
			privateKeyFile: "",
			certFile:       "",
			wantErr:        nil,
		},
		{
			name:           "cert filename absent",
			privateKeyFile: "keyfile",
			certFile:       "",
			wantErr:        fmt.Errorf("private key file provided but certificate file missing"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := server.validateFiles(tt.privateKeyFile, tt.certFile)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				if tt.wantErr != nil && gotErr != nil && tt.wantErr.Error() != gotErr.Error() {
					t.Errorf("error got: '%v', want: '%v'", gotErr, tt.wantErr)
				}
			}
		})
	}
}
