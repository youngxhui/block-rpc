package client_test

import (
	"net/http"
	"testing"

	"github.com/youngxhui/block-rpc/client"
)

func TestNewClient(t *testing.T) {
	go func() {
		http.ListenAndServe("0.0.0.0:8601", nil)
	}()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		opts    []client.Option
		want    client.Client
		wantErr bool
	}{
		{name: "Localhost", opts: []client.Option{
			client.WithAddress("localhost:8601"),
		}, want: client.BlockClient{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := client.NewClient(tt.opts...)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewClient() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewClient() succeeded unexpectedly")
			}
			if got == nil {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
