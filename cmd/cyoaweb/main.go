package main

import (
	"cyoa/pkg"
	"flag"
	"fmt"
	"os"
)

func main() {
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

	fmt.Print("JSON:")
	fmt.Printf("%+v", story)
}