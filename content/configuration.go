package content

import (
	"github.com/yannlandry/simple-markdown-website/util"
)

type NavigationConfiguration struct {
	URL   string `yaml:"URL"`
	Title string `yaml:"Title"`
}

type PagesConfiguration struct {
	Title string `yaml:"Title"`
	Path  string `yaml:"Path"`
}

type TemplatesConfiguration struct {
	Base  string `yaml:"Base"`
	Page  string `yaml:"Page"`
	Error string `yaml:"Error"`
}

type Configuration struct {
	Navigation []*NavigationConfiguration     `yaml:"Navigation"`
	Pages      map[string]*PagesConfiguration `yaml:"Pages"`
	Proxies    map[string]string              `yaml:"Proxies"`
	Templates  TemplatesConfiguration         `yaml:"Templates"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Pages: map[string]*PagesConfiguration{},
	}
}

func (this *Configuration) Load(path *util.Path) error {
	return util.LoadYAML(path.With("configuration.yaml"), this)
}
