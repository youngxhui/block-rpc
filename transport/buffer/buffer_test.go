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
		{name: "size 1", data: []byte{'1'}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := buffer.NewByteBuffer()
			i, err := b.Write(tt.data)
			if err != nil {
				t.Errorf("Write Buffer Error %v", err)
			}
			t.Logf("Buffer Write Len %d", i)
			got := b.Len()
			if true {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
