package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/lobre/rM/rm"
)

func main() {
	// Get filename from flags
	var file = flag.String("file", "", "provide a zip file")
	flag.Parse()
	if *file == "" {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	notebook, err := rm.NewNotebook(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spew.Dump(notebook)
}
