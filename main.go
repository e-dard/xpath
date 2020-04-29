package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antchfx/htmlquery"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Not enough arguments")
	}

	doc, err := htmlquery.LoadURL(args[1])
	if err != nil {
		log.Fatalf("Failed to load URL: %v", err)
	}

	nodes, err := htmlquery.QueryAll(doc, args[2])
	if err != nil {
		log.Fatalf(`Invalid xpath expression: %q`, args[2])
	}

	fmt.Printf("%s value=%d\n", args[3], len(nodes))
}
