package codec

type PayloadCodec interface {
	Encode(v any) ([]byte, error)
	Decode(data []byte, v any) error
}
