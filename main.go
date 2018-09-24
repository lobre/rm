package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Get filename from flags
	var file = flag.String("file", "", "provide a lines file")
	flag.Parse()
	if *file == "" {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	notebook := parse(*file)
	fmt.Printf("%#v", notebook)
}
