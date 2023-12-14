package functions

import (
	"net/url"
	"reflect"
	"testing"
)

const (
	bucket = "bucket"
	key    = "key"
)

func TestParseS3URI(t *testing.T) {
	table := []struct {
		wantErr     error
		name        string
		templateURL string
	}{
		{
			name:        "https host style",
			templateURL: "http://bucket.s3-region.amazonaws.com/key",
			wantErr:     nil,
		}, {
			name:        "https path style",
			templateURL: "http://s3-region.amazonaws.com/bucket/key",
			wantErr:     nil,
		}, {
			name:        "dualstack 1",
			templateURL: "https://s3.dualstack.region.amazonaws.com/bucket/key",
			wantErr:     nil,
		}, {
			name:        "dualstack 2",
			templateURL: "http://bucket.s3.dualstack.region.amazonaws.com/key",
			wantErr:     nil,
		}, {
			name:        "static 1",
			templateURL: "http://bucket.s3-website.region.amazonaws.com/key",
			wantErr:     nil,
		}, {
			name:        "static 2",
			templateURL: "http://bucket.s3-website-region.amazonaws.com/key",
			wantErr:     nil,
		}, {
			name:        "s3 1",
			templateURL: "https://s3.region.amazonaws.com/bucket/key",
			wantErr:     nil,
		}, {
			name:        "s3 2",
			templateURL: "http://s3-region.amazonaws.com/bucket/key",
			wantErr:     nil,
		}, {
			name:        "s3 3",
			templateURL: "https://s3.amazonaws.com/bucket/key",
			wantErr:     nil,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.templateURL)
			s3u, err := ParseS3URI(u)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("unexpected error; got: '%+v'", reflect.TypeOf(err))
				}
			} else {
				if *s3u.Bucket != bucket || *s3u.Key != key {
					t.Errorf("unexpected metadata; got '%+v'", s3u)
				}
			}
		})
	}
}
