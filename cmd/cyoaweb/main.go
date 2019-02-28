package main

import (
	"cyoa/pkg"
	"encoding/json"
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


	d := json.NewDecoder(file)
	var story pkg.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Print("JSON:")
	fmt.Printf("%+v", story)
}