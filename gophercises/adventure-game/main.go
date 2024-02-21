package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//port := flag.Int("port", 8080, "port to listen on")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
