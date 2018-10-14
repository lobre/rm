package notebook_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/lobre/rm/notebook"
)

func TestDecode(t *testing.T) {
	// Open file
	f, err := os.Open("../examples/Test.zip")
	if err != nil {
		t.Errorf("Can't open file")
	}

	// Get file stats
	fi, err := f.Stat()
	if err != nil {
		t.Errorf("Can't obtain file stat")
	}

	// Decode zip notebook
	n := notebook.New()
	err = n.Decode(f, fi.Size(), "Test")
	if err != nil {
		t.Errorf("Impossible to parse notebook: %v", err)
	}

	// Test notebook name
	if n.Name != "Test" {
		t.Errorf("Wrong notebook name")
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
}

func ExampleNew() {
	// Open file
	f, err := os.Open("../examples/Test.zip")
	if err != nil {
		log.Fatal(err)
	}

	// Get file stats
	fi, err := f.Stat()
	if err != nil {
	}

	// Decode zip notebook
	n := notebook.New()
	err = n.Decode(f, fi.Size(), "Test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n.Name)
	// Output:
	// Test
}
