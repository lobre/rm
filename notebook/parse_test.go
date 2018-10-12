package notebook

import (
	"fmt"
	"testing"
)

func TestNewNotebook(t *testing.T) {
	// Parse notebook
	n, err := NewNotebook("../examples/Test.zip")
	if err != nil {
		t.Errorf("Impossible to parse notebook: %v", err)
	}

	// Test notebook name
	if n.Name != "Test" {
		t.Errorf("Wrong notebook name")
	}

	// Test notebook hash
	if n.Hash != "f10aa4b9d09e43c0765e8284a4f9bb65" {
		t.Errorf("Notebook hash is not correct")
	}

	// Test number of pages
	if len(n.Pages) != 3 {
		t.Errorf("Wrong number of pages")
	}

	// Test notebook content file
	if n.Content.ExtraMetadata.LastTool != "SharpPencil" {
		t.Errorf("Wrong LastTool read")
	}

	// Test pagedata
	if n.Pages[0].Template != "Blank" {
		t.Errorf("Wrong template for page one")
	}

	// Test pdf extract
	if err := n.ExtractPdf("../examples/Test.pdf"); err != nil {
		t.Errorf("Cannot extract pdf: %v", err)
	}

	// Test thumbnail extract
	if err := n.ExtractThumbnail(0, "../examples/Thumb.jpg"); err != nil {
		t.Errorf("Cannot extract thumbnail: %v", err)
	}
}

func ExampleNewNotebook() {
	fileName := "../examples/Test.zip"
	notebook, _ := NewNotebook(fileName)
	fmt.Println(notebook.Name)
	// Output:
	// Test
}
