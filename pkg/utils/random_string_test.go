package utils

import (
	"testing"
)

func TestGenRandomString(t *testing.T) {

	table := []struct {
		name string
		want int
	}{
		{
			name: "gen random string 1",
			want: 3,
		},
		{
			name: "gen random string 2",
			want: 6,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got := GenRandomString(tt.want)
			if len(got) != tt.want {
				t.Errorf("got: '%v', want: '%v'", len(got), tt.want)
			}
		})
	}
}
