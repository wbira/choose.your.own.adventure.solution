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

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	story Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]


	if chapter, ok := h.story[path]; ok {
		err := tpl.Execute(w, chapter)
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
