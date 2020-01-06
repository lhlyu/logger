package logger

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Formatter interface {
	Format(v interface{}) ([]byte, error) // 格式化
	NewLine() bool                        // 是否要换行
}

var (
	TEXT        Formatter = new(textFormatter)
	JSON        Formatter = new(jsonFormatter)
	JSON_INDENT Formatter = new(jsonIndentFormatter)
	XML         Formatter = new(xmlFormatter)
	XML_INDENT  Formatter = new(xmlIndentFormatter)
	YAML        Formatter = new(yamlFormatter)
)

// text format
type textFormatter struct{}

func (textFormatter) Format(v interface{}) ([]byte, error) {
	if b, ok := v.([]byte); ok {
		return b, nil
	}
	if s, ok := v.(string); ok {
		return []byte(s), nil
	}
	return []byte(fmt.Sprint(v)), errors.New("this is not string")
}

func (textFormatter) NewLine() bool {
	return false
}

// json format
type jsonFormatter struct{}

func (jsonFormatter) Format(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonFormatter) NewLine() bool {
	return false
}

// json indent format
type jsonIndentFormatter struct{}

func (jsonIndentFormatter) Format(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (jsonIndentFormatter) NewLine() bool {
	return true
}

// xml format
type xmlFormatter struct{}

func (xmlFormatter) Format(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (xmlFormatter) NewLine() bool {
	return false
}

// xml indent format
type xmlIndentFormatter struct{}

func (xmlIndentFormatter) Format(v interface{}) ([]byte, error) {
	return xml.MarshalIndent(v, "", "  ")
}

func (xmlIndentFormatter) NewLine() bool {
	return true
}

// yaml format
type yamlFormatter struct{}

func (yamlFormatter) Format(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (yamlFormatter) NewLine() bool {
	return true
}
