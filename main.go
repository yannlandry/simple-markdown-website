package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/simple-markdown-website/content"
	"github.com/yannlandry/simple-markdown-website/handler"
	"github.com/yannlandry/simple-markdown-website/util"
)

func main() {
	log.Println("Starting Simple Markdown Website...")

	// Command-line arguments.
	var rootRaw string
	flag.StringVar(&rootRaw, "root", "", "Path to the directory defining the website's content.")
	var baseURLRaw string
	flag.StringVar(&baseURLRaw, "base-url", "", "URL of the website to be prepended to links.")
	var staticURLRaw string
	flag.StringVar(&staticURLRaw, "static-url", "", "URL of static assets to be prepended to images, stylesheets, etc.")
	var port int
	flag.IntVar(&port, "port", 8080, "Port on which the Golang app should listen.")
	flag.Parse()

	// Check command-line arguments.
	if rootRaw == "" {
		log.Fatalln("`--root` is a required argument.")
	}
	if baseURLRaw == "" {
		log.Fatalln("`--base-url` is a required argument.")
	}
	if staticURLRaw == "" {
		log.Fatalln("`--static-url` is a required argument.")
	}

	// Convert the content path to a smart path.
	root := util.NewPath(rootRaw)

	// Instantiate `URLBuilder`s.
	var err error
	var baseURL *util.URLBuilder
	var staticURL *util.URLBuilder
	if baseURL, err = util.NewURLBuilder(baseURLRaw); err != nil {
		log.Fatalf("Failed parsing the base URL: %s\n", err)
	}
	if staticURL, err = util.NewURLBuilder(staticURLRaw); err != nil {
		log.Fatalf("Failed parsing the static URL: %s\n", err)
	}

	// Instantiate markdown renderer
	util.Markdown = util.NewMarkdownEngine(baseURL, staticURL)

	// Load configuration.
	config := content.NewConfiguration()
	if err := config.Load(root); err != nil {
		log.Fatalf("Failed loading the website configuration: %s\n", err)
	}
	log.Println("Done loading configuration.")

	// Load pages.
	pages := content.NewPages()
	if err := pages.Load(root, config); err != nil {
		log.Fatalf("Failed loading pages: %s\n", err)
	}
	log.Println("Done loading pages.")

	// Load base template and create builder.
	builder := util.NewTemplateBuilder(root.With(config.Templates.Base))
	builder.BaseURL = baseURL
	builder.StaticURL = staticURL

	// Handler factory.
	handlers := handler.NewHandlerFactory(root, config, pages, builder)

	// Router.
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Page())
	router.HandleFunc("/{slug}", handlers.Page())
	router.HandleFunc("/{slug}/", handlers.Page())

	// Server.
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
	}
	log.Println(server.ListenAndServe())
}
