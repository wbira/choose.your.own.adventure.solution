package main

import (
	"cyoa/pkg"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "Port on which web app is running on")
	filename := flag.String("file", "gopher.json", "Path to json that contains story")
	flag.Parse()
	fmt.Printf("Using story from file %v\n", *filename)

	file, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := pkg.JsonStory(file)
	if err != nil {
		panic(err)
	}
	tpl := template.Must(template.New("").Parse(styledTemplates))
	httpHandler := pkg.NewHandler(story, pkg.WithTemplate(tpl))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), httpHandler))
}


var styledTemplates = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
