package handler

import (
	"html/template"
	"net/http"
)

type PagePresenter struct {
	Content template.HTML
}

func (this *HandlerFactory) Page(slug string) func(http.ResponseWriter, *http.Request) {
	// Load the template for this handler.
	template, _ := this.builder.Load(this.root.With(this.configuration.Templates.Page))
	// Generate handler function using the template, configuration, etc.
	return func(response http.ResponseWriter, request *http.Request) {
		// Load configuration for this specific page.
		configuration, ok := this.configuration.Pages[slug]
		if !ok {
			Error(this.errors, this.configuration, 404, response, request)
			return
		}
		// Load Markdown-formatted content for this specific page.
		content, ok := this.pages.Content[slug]
		if !ok {
			Error(this.errors, this.configuration, 500, response, request)
			return
		}
		// Fill in the presenter and execute template.
		presenter := NewBasePresenter(this.configuration, &PagePresenter{
			Content: content,
		})
		presenter.WindowTitle = configuration.Title
		ExecuteTemplate(template, response, presenter)
	}
}
