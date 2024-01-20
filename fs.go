package main

import (
	_ "embed"
	"strings"

	"bytes"
	"text/template"
)

//go:embed tpl.txt
var fs string

type (
	Option struct {
		ClassName        string
		DartMapKeyType   string
		DartMapValueType string
		Fields           []*Field
	}
	Field struct {
		JsonName      string
		FieldName     string
		DataType      string
		InnerDataType string
		Default       string
		Array         bool
		ClassArray    bool
		Object        bool
	}
)

func executeTpl(options []*Option) string {
	var (
		buf     = &bytes.Buffer{}
		tpl     = template.Must(template.New("").Parse(fs))
		err     error
		results []string
	)

	for _, opt := range options {
		buf.Reset()
		if err = tpl.Execute(buf, &opt); err != nil {
			panic(err.Error())
		}
		results = append(results, buf.String())
	}

	return strings.Join(results, "\n")
}
