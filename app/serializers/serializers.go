package serializers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
)

// default serializer is json encoder
var DefaultSerializer JSONEncoder

// Serializer provides an interface for providing custom serializers for values.
type Serializer interface {
	Serialize(src interface{}) ([]byte, error)
	Deserialize(src []byte, dst interface{}) error
}

// GobEncoder encodes values using encoding/gob. This is the simplest
// encoder and can handle complex types via gob.Register.
type GobEncoder struct{}

// JSONEncoder encodes cookie values using encoding/json. Users who wish to
// encode complex types need to satisfy the json.Marshaller and
// json.Unmarshaller interfaces.
type JSONEncoder struct{}

// NopEncoder does not encode cookie values, and instead simply accepts a []byte
// (as an interface{}) and returns a []byte. This is particularly useful when
// you encoding an object upstream and do not wish to re-encode it.
type NopEncoder struct{}

// imitate cookie sz

type UnSecureCookieImitater struct{}

func (s *UnSecureCookieImitater) Decode(name, value string, dst interface{}) error {
	v, ok := dst.(*string)
	if !ok {
		return errors.New("convert error")
	}

	*v = value
	return nil
}

func (s *UnSecureCookieImitater) Encode(name string, value interface{}) (string, error) {
	res, ok := value.(string)
	if !ok {
		return "", nil
	}
	return res, nil
}

// Serialization --------------------------------------------------------------

// Serialize encodes a value using gob.
func (e GobEncoder) Serialize(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, cookieError{cause: err, typ: usageError}
	}
	return buf.Bytes(), nil
}

// Deserialize decodes a value using gob.
func (e GobEncoder) Deserialize(src []byte, dst interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(src))
	if err := dec.Decode(dst); err != nil {
		return cookieError{cause: err, typ: decodeError}
	}
	return nil
}

// Serialize encodes a value using encoding/json.
func (e JSONEncoder) Serialize(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, cookieError{cause: err, typ: usageError}
	}
	return buf.Bytes(), nil
}

// Deserialize decodes a value using encoding/json.
func (e JSONEncoder) Deserialize(src []byte, dst interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(dst); err != nil {
		return cookieError{cause: err, typ: decodeError}
	}
	return nil
}

// Serialize passes a []byte through as-is.
func (e NopEncoder) Serialize(src interface{}) ([]byte, error) {
	if b, ok := src.([]byte); ok {
		return b, nil
	}

	return nil, errValueNotByte
}

// Deserialize passes a []byte through as-is.
func (e NopEncoder) Deserialize(src []byte, dst interface{}) error {
	if dat, ok := dst.(*[]byte); ok {
		*dat = src
		return nil
	}
	return errValueNotBytePtr
}
