package rm

import (
	"fmt"
	"testing"
)

func TestNewNotebook(t *testing.T) {
	// Parse notebook
	n, err := NewNotebook("examples/Test.zip")
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

	// Test notebook content file
	if n.Content.ExtraMetadata.LastTool != "SharpPencil" {
		t.Errorf("Wrong LastTool read")
	}
}

func ExampleNewNotebook() {
	fileName := "examples/Test.zip"
	notebook, _ := NewNotebook(fileName)
	fmt.Println(notebook.Name)
	// Output:
	// Test
}
