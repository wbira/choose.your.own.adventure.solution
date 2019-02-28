package main

import (
	"cyoa/pkg"
	"flag"
	"fmt"
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

	httpHandler := pkg.NewHandler(story)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), httpHandler))
}
