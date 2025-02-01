package buffer_test

import (
	"testing"

	"github.com/youngxhui/block-rpc/transport/buffer"
)

func TestByteBuffer_Len(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		data []byte
		want int
	}{
		{name: "buffer is empty", data: []byte{}, want: int(0)},
		{name: "size 1", data: []byte{'1'}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buffer.NewByteBuffer()
			_, err := b.Write(tt.data)
			if err != nil {
				t.Errorf("Write Buffer Error %v", err)
			}
			got := b.Len()
			if got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_Read(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n       int
		data    []byte
		want    []byte
		wantErr bool
	}{
		{name: "read empty", n: 0, data: []byte("hello"), want: []byte{}, wantErr: false},
		{name: "read 1 size", n: 1, data: []byte("hello"), want: []byte("h"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buffer.NewByteBuffer()
			b.Write(tt.data)
			got, gotErr := b.Read(tt.n)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Read() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Read() succeeded unexpectedly")
			}
			if len(got) != len(tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
