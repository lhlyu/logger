package pio

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

type Marshaled interface {
	Marshal() ([]byte, error)
}

func fromMarshaled(self Marshaled) Marshaler {
	return MarshalerFunc(func(v interface{}) ([]byte, error) {
		return self.Marshal()
	})
}

type MarshalerFunc func(v interface{}) ([]byte, error)

func (m MarshalerFunc) Marshal(v interface{}) ([]byte, error) {
	return m(v)
}

var ErrMarshalNotResponsible = errors.New("this marshaler is not responsible for this type of data")

var ErrMarshalNotFound = errors.New("no marshaler found for this type of dat")

var Text = MarshalerFunc(func(v interface{}) ([]byte, error) {
	if b, ok := v.([]byte); ok {
		return b, nil
	}
	if s, ok := v.(string); ok {
		return []byte(s), nil
	}

	return nil, ErrMarshalNotResponsible
})

var (
	JSON       = MarshalerFunc(json.Marshal)
	JSONIndent = MarshalerFunc(func(v interface{}) ([]byte, error) {
		return json.MarshalIndent(v, "", "  ")
	})
	XML       = MarshalerFunc(xml.Marshal)
	XMLIndent = MarshalerFunc(func(v interface{}) ([]byte, error) {
		return xml.MarshalIndent(v, "", " ")
	})
)
