package json

import (
	encoding "encoding/json"
	"io"
)

// Decode ...
func Decode(v interface{}, r io.Reader) error {
	return encoding.NewDecoder(r).Decode(v)
}

// DecodeString ...
func DecodeString(v interface{}, s string) error {
	return encoding.Unmarshal([]byte(s), v)
}

// EncodeString ...
func EncodeString(v interface{}) (string, error) {
	data, err := encoding.Marshal(v)
	return string(data), err
}
