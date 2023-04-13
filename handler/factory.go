package handler

import (
	"html/template"

	"github.com/yannlandry/simple-markdown-website/content"
	"github.com/yannlandry/simple-markdown-website/util"
)

type HandlerFactory struct {
	root          *util.Path
	configuration *content.Configuration
	pages         *content.Pages
	builder       *util.TemplateBuilder
	errors        *template.Template
}

func NewHandlerFactory(root *util.Path, configuration *content.Configuration, pages *content.Pages, builder *util.TemplateBuilder) *HandlerFactory {
	// Load template for error pages.
	errors, _ := builder.Load(root.With(configuration.Templates.Error))
	// Create handler factory.
	return &HandlerFactory{
		root:          root,
		configuration: configuration,
		pages:         pages,
		builder:       builder,
		errors:        errors,
	}
}
