package envconfig

import (
	"encoding"
	"errors"
	"reflect"
	"regexp"
)

var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

var gatherRegexp = regexp.MustCompile("([^A-Z]+|[A-Z]+[^A-Z]+|[A-Z]+)")
var acronymRegexp = regexp.MustCompile("([A-Z]+)([A-Z][^A-Z]+)")

type ParseError struct {
	KeyName   string
	FieldName string
	TypeName  string
	Value     string
	Err       error
}

type Decoder interface {
	Decode(value string) error
}

type Setter interface {
	Set(value string) error
}

func (e *ParseError) Error() string { _ = "STUB: not implemented"; return "" }

type varInfo struct {
	Name  string
	Alt   string
	Key   string
	Field reflect.Value
	Tags  reflect.StructTag
}

func gatherInfo(prefix string, spec interface{}) ([]varInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CheckDisallowed(prefix string, spec interface{}) error { _ = "STUB: not implemented"; return nil }

func Process(prefix string, spec interface{}) error { _ = "STUB: not implemented"; return nil }

func MustProcess(prefix string, spec interface{}) { _ = "STUB: not implemented"; return }

func processField(value string, field reflect.Value) error { _ = "STUB: not implemented"; return nil }

func interfaceFrom(field reflect.Value, fn func(interface{}, *bool)) {
	_ = "STUB: not implemented"
	return
}

func decoderFrom(field reflect.Value) (d Decoder) { _ = "STUB: not implemented"; return *new(Decoder) }

func setterFrom(field reflect.Value) (s Setter) { _ = "STUB: not implemented"; return *new(Setter) }

func textUnmarshaler(field reflect.Value) (t encoding.TextUnmarshaler) {
	_ = "STUB: not implemented"
	return *new(encoding.TextUnmarshaler)
}

func binaryUnmarshaler(field reflect.Value) (b encoding.BinaryUnmarshaler) {
	_ = "STUB: not implemented"
	return *new(encoding.BinaryUnmarshaler)
}

func isTrue(s string) bool { _ = "STUB: not implemented"; return false }
