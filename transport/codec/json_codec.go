package codec

import (
	"encoding/json"
	"errors"
)

type JSONCodec struct{}

func (c *JSONCodec) Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JSONCodec) Decode(data []byte, v any) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}
	return json.Unmarshal(data, v)
}
