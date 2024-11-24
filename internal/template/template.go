package template

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/zygote-sh/zygote/internal/context"
)

type Template string

func (t Template) Parse() Template {
	if value, err := parse(string(t), context.Current); err == nil {
		return Template(value)
	}
	return t
}

func (t Template) String() string {
	return string(t)
}

func parse(text string, ctx interface{}) (string, error) {
	if !strings.Contains(text, "{{") || !strings.Contains(text, "}}") {
		return text, nil
	}

	parsedTemplate, err := template.New("alias").Funcs(funcMap()).Parse(text)
	if err != nil {
		println(err.Error())
		return "", err
	}

	buffer := new(bytes.Buffer)
	defer buffer.Reset()

	err = parsedTemplate.Execute(buffer, ctx)
	if err != nil {
		println(err.Error())
		return "", err
	}

	return buffer.String(), nil
}

type Templates []Template

func (t Templates) Strings() []string {
	var s []string
	for _, v := range t {
		s = append(s, string(v))
	}
	return s
}

func (t Templates) Parse() []string {
	var s []string
	for _, v := range t {
		s = append(s, v.Parse().String())
	}
	return s
}
