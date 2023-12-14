

package helmv3

import (
	"reflect"
	"testing"

	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
)

func TestLoadIacFile(t *testing.T) {

	table := []struct {
		name     string
		filePath string
		options  map[string]interface{}
		helmv3   HelmV3
		typeOnly bool
		want     output.AllResourceConfigs
		wantErr  error
	}{
		{
			name:     "load iac file is not supported for helm",
			filePath: "/dummyfilepath.yaml",
			helmv3:   HelmV3{},
			wantErr:  errLoadIacFileNotSupported,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := tt.helmv3.LoadIacFile(tt.filePath, tt.options)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", gotErr, tt.wantErr)
			} else if tt.typeOnly && (reflect.TypeOf(gotErr)) != reflect.TypeOf(tt.wantErr) {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", reflect.TypeOf(gotErr), reflect.TypeOf(tt.wantErr))
			}
		})
	}

}
