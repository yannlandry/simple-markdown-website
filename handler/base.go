package handler

import (
	"github.com/yannlandry/simple-markdown-website/content"
)

type BasePresenter struct {
	WindowTitle   string
	Configuration *content.Configuration
	Presenter     interface{}
}

func NewBasePresenter(configuration *content.Configuration, presenter interface{}) *BasePresenter {
	return &BasePresenter{
		Configuration: configuration,
		Presenter:     presenter,
	}
}
