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
	if n.hash != "01fb883ad1d03543b44ae43a4ce23d51" {
		t.Errorf("Notebook hash is not correct")
	}

	// Test notebook content file
	if n.content.ExtraMetadata.LastTool != "Fineliner" {
		t.Errorf("Wrong content lasttool read")
	}
}

func ExampleNewNotebook() {
	fileName := "examples/Test.zip"
	notebook, _ := NewNotebook(fileName)
	fmt.Println(notebook.Name)
	// Output:
	// Test
}
