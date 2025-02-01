package codec_test

import (
	"testing"

	"github.com/youngxhui/block-rpc/transport/codec"
)

type demo struct {
	A string `json:"a"`
}

func TestJSONCodec_Encode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a       string
		want    []byte
		wantErr bool
	}{
		{name: "json empty", a: "", want: []byte{'{', '"', 'a', '"', ':', '"', '"', '}'}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c codec.JSONCodec
			var d demo
			d.A = tt.a

			got, gotErr := c.Encode(d)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Encode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Encode() succeeded unexpectedly")
			}
			if len(tt.want) != len(got) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONCodec_Decode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a       string
		data    []byte
		wantErr bool
	}{
		{name: "not error", data: []byte{'{', '"', 'a', '"', ':', '"', 'a', '"', '}'}, a: "a", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c codec.JSONCodec
			d := demo{
				A: tt.a,
			}
			gotErr := c.Decode(tt.data, &d)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Decode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Decode() succeeded unexpectedly")
			}
		})
	}
}
