package serializers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
)

// default serializer is json encoder
var DefaultSerializer Serializer

// Serializer provides an interface for providing custom serializers for values.
type Serializer interface {
	Serialize(src interface{}) ([]byte, error)
	Deserialize(src []byte, dst interface{}) error
}

// GobEncoder encodes values using encoding/gob. This is the simplest
// encoder and can handle complex types via gob.Register.
type GobEncoder struct{}

// JSONEncoder encodes values using encoding/json. Users who wish to
// encode complex types need to satisfy the json.Marshaller and
// json.Unmarshaller interfaces.
type JSONEncoder struct{}

// NopEncoder does not encode values, and instead simply accepts a []byte
// (as an interface{}) and returns a []byte. This is particularly useful when
// you encoding an object upstream and do not wish to re-encode it.
type NopEncoder struct{}

// Serialization --------------------------------------------------------------

// Serialize encodes a value using gob.
func (e *GobEncoder) Serialize(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, errors.New("encode error")
	}
	return buf.Bytes(), nil
}

// Deserialize decodes a value using gob.
func (e *GobEncoder) Deserialize(src []byte, dst interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(src))
	if err := dec.Decode(dst); err != nil {
		return errors.New("decode error")
	}
	return nil
}

// Serialize encodes a value using encoding/json.
func (e *JSONEncoder) Serialize(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, errors.New("encode error")
	}
	return buf.Bytes(), nil
}

// Deserialize decodes a value using encoding/json.
func (e *JSONEncoder) Deserialize(src []byte, dst interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(dst); err != nil {
		return errors.New("decode error")
	}
	return nil
}

// Serialize passes a []byte through as-is.
func (e *NopEncoder) Serialize(src interface{}) ([]byte, error) {
	if b, ok := src.([]byte); ok {
		return b, nil
	}

	return nil, errors.New("value not byte")
}

// Deserialize passes a []byte through as-is.
func (e *NopEncoder) Deserialize(src []byte, dst interface{}) error {
	if dat, ok := dst.(*[]byte); ok {
		*dat = src
		return nil
	}
	return errors.New("value not byte ptr")
}

// Other Function --------------------------------------------------------------
