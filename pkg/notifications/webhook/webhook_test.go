

package webhook

import "testing"

func TestWebhookInit(t *testing.T) {
	testURL := "testURL"
	testToken := "testToken"

	type args struct {
		config interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		assert  bool
		url     string
		token   string
	}{
		{
			name: "nil config",
			args: args{
				config: nil,
			},
			wantErr: true,
		},
		{
			name: "valid webhook config data",
			args: args{
				config: map[string]interface{}{
					"url":   testURL,
					"token": testToken,
				},
			},
			assert: true,
			url:    testURL,
			token:  testToken,
		},
		{
			name: "invalid webhook config data",
			args: args{
				config: struct {
					url   string
					token string
				}{
					url:   testURL,
					token: testToken,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{}
			if err := w.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Webhook.Init() got error: %v, wantErr: %v", err, tt.wantErr)
			}
			if tt.assert {
				if w.URL != tt.url {
					t.Errorf("Webhook.Init() got url: %v, want url: %v", w.URL, tt.url)
				}

				if w.Token != tt.token {
					t.Errorf("Webhook.Init() got token: %v, want token: %v", w.Token, tt.token)
				}
			}
		})
	}
}
