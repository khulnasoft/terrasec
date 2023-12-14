package runtime

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestValidateInputs(t *testing.T) {

	table := []struct {
		name     string
		executor Executor
		wantErr  error
	}{
		{
			name: "valid filePath",
			executor: Executor{
				filePath:    filepath.Join(testDataDir, "testfile"),
				dirPath:     "",
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
			},
			wantErr: nil,
		},
		{
			name: "valid dirPath",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
			},
			wantErr: nil,
		},
		{
			name: "valid filePath",
			executor: Executor{
				filePath:    filepath.Join(testDataDir, "testfile"),
				dirPath:     "",
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
				severity:    "high",
			},
			wantErr: nil,
		},
		{
			name: "valid dirPath",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
				severity:    "MEDIUM",
			},
			wantErr: nil,
		},
		{
			name: "valid dirPath",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v12",
				severity:    " LOW ",
				categories:  []string{" identity And ACCESS Management ", "data Protection "},
			},
			wantErr: nil,
		},
		{
			name: "empty iac path",
			executor: Executor{
				filePath: "",
				dirPath:  "",
			},
			wantErr: errEmptyIacPath,
		},
		{
			name: "filepath does not exist",
			executor: Executor{
				filePath: filepath.Join(testDataDir, "notthere"),
			},
			wantErr: errFileNotExists,
		},
		{
			name: "directory does not exist",
			executor: Executor{
				dirPath: filepath.Join(testDataDir, "notthere"),
			},
			wantErr: errDirNotExists,
		},
		{
			// should error out in validations if -f option is not a file
			name: "valid directory passed as file path",
			executor: Executor{
				filePath: testDir,
			},
			wantErr: errNotValidFile,
		},
		{
			// should error out in validations if -d option is not a dir
			name: "valid directory passed as file path",
			executor: Executor{
				dirPath: filepath.Join(testDir, "testfile"),
			},
			wantErr: errNotValidDir,
		},
		{
			name: "invalid iac type",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "notthere",
				iacVersion:  "v14",
			},
			wantErr: errIacNotSupported,
		},
		{
			name: "invalid iac version",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "notthere",
			},
			wantErr: errIacNotSupported,
		},
		{
			name: "invalid severity",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
				severity:    "HGIH",
			},
			wantErr: errSeverityNotSupported,
		},
		{
			name: "invalid category",
			executor: Executor{
				filePath:    "",
				dirPath:     testDir,
				policyTypes: []string{"aws"},
				iacType:     "terraform",
				iacVersion:  "v14",
				severity:    "HGIH",
				categories:  []string{"DTA PROTECTIO"},
			},
			wantErr: fmt.Errorf(errCategoryNotSupported, []string{"DTA PROTECTIO"}),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.executor.ValidateInputs()
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("unexpected error, gotErr: '%v', wantErr: '%v'", gotErr, tt.wantErr)
			}
		})
	}
}
