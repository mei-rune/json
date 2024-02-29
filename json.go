package json

import (
	// "encoding/json"
	"bytes"
	"io"

	json "github.com/goccy/go-json"
)

type Number = json.Number
type Unmarshaler = json.Unmarshaler
type UnmarshalTypeError = json.UnmarshalTypeError
type UnmarshalFieldError = json.UnmarshalFieldError
type InvalidUnmarshalError = json.InvalidUnmarshalError

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type Marshaler = json.Marshaler
type UnsupportedTypeError = json.UnsupportedTypeError
type UnsupportedValueError = json.UnsupportedValueError
type InvalidUTF8Error = json.InvalidUTF8Error
type MarshalerError = json.MarshalerError

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func HTMLEscape(dst *bytes.Buffer, src []byte) {
	json.HTMLEscape(dst, src)
}

func Compact(dst *bytes.Buffer, src []byte) error {
	return json.Compact(dst, src)
}

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

type SyntaxError = json.SyntaxError

func Valid(data []byte) bool {
	return json.Valid(data)
}

type Decoder = json.Decoder

func NewDecoder(r io.Reader) *Decoder {
	return json.NewDecoder(r)
}

type Encoder = json.Encoder

func NewEncoder(w io.Writer) *Encoder {
	return json.NewEncoder(w)
}

// func Fuzz(data []byte) (score int) {
// 	return json.Fuzz(data)
// }
