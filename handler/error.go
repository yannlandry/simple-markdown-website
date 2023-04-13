package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/yannlandry/simple-markdown-website/content"
)

type ErrorPresenter struct {
	Code int
}

func Error(template *template.Template, configuration *content.Configuration, code int, response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(code)
	presenter := NewBasePresenter(configuration, &ErrorPresenter{
		Code: code,
	})
	presenter.WindowTitle = fmt.Sprintf("Error %d", code)
	ExecuteTemplate(template, response, presenter)
}
