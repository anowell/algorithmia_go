package main

import (
	"algorithmia"
	"fmt"
	"os"
	"strings"
)

func main() {
	api_key := os.Getenv("ALGORITHMIA_API_KEY")

	if len(api_key) == 0 || len(os.Args) < 2 {
		usage()
	}

	data := strings.NewReader(os.Args[1])
	results, err := algorithmia.Query("kenny/Factor", api_key, data)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	fmt.Printf("%s", results)
}

func usage() {
	fmt.Printf("Usage: ALGORITHMIA_API_KEY=123456 factor NUMBER")
	os.Exit(1)
}
