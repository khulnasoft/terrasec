

package tfplan

import (
	"reflect"
	"testing"
)

func TestLoadIacDir(t *testing.T) {

	t.Run("directory not supported", func(t *testing.T) {
		var (
			dirPath = "some-path"
			tfplan  = TFPlan{}
			wantErr = errIacDirNotSupport
			options = make(map[string]interface{})
		)
		_, err := tfplan.LoadIacDir(dirPath, options)
		if !reflect.DeepEqual(wantErr, err) {
			t.Errorf("error want: '%v', got: '%v'", wantErr, err)
		}
	})
}
