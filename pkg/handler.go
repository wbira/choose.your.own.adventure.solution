package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

type HandlerOptions func(r *handler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func (h *handler) {
		h.template = t
	}
}

func NewHandler(s Story, options ...HandlerOptions) http.Handler {
	h := handler{ defaultPathFactory, s, tpl}

	for _, opt := range options {
		opt(&h)
	}

	return h
}

type handler struct {
	pathFn func(r *http.Request) string
	story Story
	template *template.Template
}

func defaultPathFactory(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.story[path]; ok {
		err := h.template.Execute(w, chapter)
		if err != nil {
			fmt.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Status not found", http.StatusNotFound)
}

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose your own adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{ range .Paragraphs}}
    <p>{{.}}</p>
    {{ end }}
    <ul>
      {{range .Options}}
      <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>`
