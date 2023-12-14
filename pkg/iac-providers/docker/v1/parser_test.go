

package dockerv1

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct {
		name     string
		filePath string
		dockerv1 DockerV1
		want     []DockerConfig
		wantErr  error
	}{
		{
			name:     "valid docker file",
			filePath: filepath.Join(fileTestDataDir, "dockerfile-testparse-function"),
			dockerv1: DockerV1{},
			wantErr:  nil,
			want:     []DockerConfig{{Cmd: "from", Value: "runatlantis/atlantis:v0.16.1", Line: 1}, {Cmd: "maintainer", Value: "tenable", Line: 2}, {Cmd: "label", Value: "key \"value\"", Line: 3}, {Cmd: "workdir", Value: "test", Line: 4}, {Cmd: "env", Value: "DEFAULT_TERRASEC_VERSION 1.5.1", Line: 5}, {Cmd: "env", Value: "PLANFILE tfplan", Line: 6}, {Cmd: "add", Value: "setup.sh terrasec.sh launch-atlantis.sh entrypoint.sh /usr/local/bin/", Line: 7}, {Cmd: "run", Value: "mkdir -p /etc/atlantis/ &&     chmod +x /usr/local/bin/*.sh &&     /usr/local/bin/setup.sh", Line: 8}, {Cmd: "copy", Value: "terrasec-workflow.yaml /etc/atlantis/workflow.yaml", Line: 11}, {Cmd: "user", Value: "atlantis", Line: 13}, {Cmd: "arg", Value: "name=defaultValue", Line: 14}, {Cmd: "run", Value: "terrasec init", Line: 15}, {Cmd: "volume", Value: "/temp", Line: 16}, {Cmd: "healthcheck", Value: "--interval=30s --timeout=30s --start-period=5s --retries=3 CMD executable", Line: 17}, {Cmd: "entrypoint", Value: "/bin/bash entrypoint.sh", Line: 18}, {Cmd: "shell", Value: "cd", Line: 19}, {Cmd: "onbuild", Value: "", Line: 20}, {Cmd: "expose", Value: "9090", Line: 21}, {Cmd: "stopsignal", Value: "1", Line: 22}, {Cmd: "cmd", Value: "server", Line: 23}},
		},
		{
			name:     "invalid  docker file path",
			filePath: filepath.Join(fileTestDataDir, "dockerfile-testparse-function1"),
			dockerv1: DockerV1{},
			want:     []DockerConfig{},
			wantErr:  fmt.Errorf("open %s: no such file or directory", filepath.Join(fileTestDataDir, "dockerfile-testparse-function1")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, _, err := tt.dockerv1.Parse(tt.filePath)
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", err, tt.wantErr)
					return
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerV1.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
