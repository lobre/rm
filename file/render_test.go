package file

import (
	"fmt"

	"github.com/lobre/rm"
)

func ExampleNewDrawer() {
	// Parse notebook from zip file
	n, err := rm.NewNotebook("../examples/Test.zip")
	if err != nil {
		fmt.Printf("Impossible to create notebook from zip: %v", err)
	}

	// Create drawer object
	f := NewDrawer(n)

	// Draw png file
	err = f.DrawPng()
	if err != nil {
		fmt.Printf("Error while generating png: %v", err)
	}

	// Draw pdf file
	err = f.DrawPdf()
	if err != nil {
		fmt.Printf("Error while generating pdf: %v", err)
	}

	// Draw svg file
	err = f.DrawSvg()
	if err != nil {
		fmt.Printf("Error while generating svg: %v", err)
	}

	fmt.Println("Notebook files generated")

	// Output:
	// Notebook files generated
}
