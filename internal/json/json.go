package json

import (
	"io"

	encoder "github.com/goccy/go-json"
)

// Decode ...
func Decode(v interface{}, r io.Reader) error {
	return encoder.NewDecoder(r).Decode(v)
}

// Encode ...
func Encode(v interface{}, w io.Writer) error {
	return encoder.NewEncoder(w).Encode(v)
}

// DecodeString ...
func DecodeString(v interface{}, s string) error {
	return encoder.Unmarshal([]byte(s), v)
}

// EncodeString ...
func EncodeString(v interface{}) (string, error) {
	data, err := encoder.Marshal(v)
	return string(data), err
}
