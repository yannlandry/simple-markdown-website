package content

import (
	"fmt"
	"html/template"

	"github.com/yannlandry/simple-markdown-website/util"
)

type Pages struct {
	Content map[string]template.HTML
}

func NewPages() *Pages {
	return &Pages{
		Content: map[string]template.HTML{},
	}
}

func (this *Pages) Load(path *util.Path, configuration *Configuration) error {
	for slug, page := range configuration.Pages {
		// Load file from configuration.
		raw, err := util.LoadFile(path.With(page.Path))
		if err != nil {
			return fmt.Errorf("failed to load page content: %s", err)
		}
		// Parse markdown into HTML.
		this.Content[slug] = template.HTML(util.Markdown.Render(raw))
	}
	return nil
}