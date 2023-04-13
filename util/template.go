package util

import (
	"fmt"
	"html/template"
	"net/url"
	"path"
	"strings"
	"time"
)

var (
	lastRestart = time.Now().Unix()
)

type TemplateBuilder struct {
	base      string
	BaseURL   *URLBuilder
	StaticURL *URLBuilder
}

func NewTemplateBuilder(base string) *TemplateBuilder {
	return &TemplateBuilder{
		base:      base,
		BaseURL:   nil,
		StaticURL: nil,
	}
}

func (this *TemplateBuilder) Load(paths ...string) (*template.Template, error) {
	arguments := []string{this.base}
	arguments = append(arguments, paths...)

	functions := template.FuncMap{
		"baseURL": func(extension string) string {
			return this.BaseURL.With(extension)
		},
		"staticURL": func(extension string) string {
			return this.StaticURL.With(extension)
		},
		"concat": func(elements ...string) string {
			return strings.Join(elements, "")
		},
		"urlEscape": func(text string) string {
			return url.QueryEscape(text)
		},
		"lastRestart": func() int64 {
			return lastRestart
		},
	}

	tmp, err := template.New(path.Base(this.base)).Funcs(functions).ParseFiles(arguments...)
	if err != nil {
		return nil, fmt.Errorf("failed loading templates `%s`: %s", paths, err)
	}

	return tmp, nil
}
