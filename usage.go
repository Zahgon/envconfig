package envconfig

import (
	"encoding"
	"io"
	"reflect"
	"text/template"
)

const (
	DefaultListFormat = `This application is configured via the environment. The following environment
variables can be used:
{{range .}}
{{usage_key .}}
  [description] {{usage_description .}}
  [type]        {{usage_type .}}
  [default]     {{usage_default .}}
  [required]    {{usage_required .}}{{end}}
`

	DefaultTableFormat = `This application is configured via the environment. The following environment
variables can be used:

KEY	TYPE	DEFAULT	REQUIRED	DESCRIPTION
{{range .}}{{usage_key .}}	{{usage_type .}}	{{usage_default .}}	{{usage_required .}}	{{usage_description .}}
{{end}}`
)

var (
	decoderType           = reflect.TypeOf((*Decoder)(nil)).Elem()
	setterType            = reflect.TypeOf((*Setter)(nil)).Elem()
	textUnmarshalerType   = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
	binaryUnmarshalerType = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

func implementsInterface(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func toTypeDescription(t reflect.Type) string { _ = "STUB: not implemented"; return "" }

func Usage(prefix string, spec interface{}) error { _ = "STUB: not implemented"; return nil }

func Usagef(prefix string, spec interface{}, out io.Writer, format string) error {
	_ = "STUB: not implemented"
	return nil
}

func Usaget(prefix string, spec interface{}, out io.Writer, tmpl *template.Template) error {
	_ = "STUB: not implemented"
	return nil
}
