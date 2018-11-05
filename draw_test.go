package notebook_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/lobre/rm/notebook"
)

func TestDrawPng(t *testing.T) {
	// Open file
	f, err := os.Open("examples/Test.zip")
	if err != nil {
		t.Error(err)
	}

	// Get file stats
	fi, err := f.Stat()
	if err != nil {
	}

	// Decode zip notebook
	n := notebook.New("Test")
	err = n.Decode(f, fi.Size())
	if err != nil {
		t.Error(err)
	}

	// Draw note
	b, err := n.DrawPng()
	if err != nil {
		t.Error(err)
	}

	// Write png to file
	if err = ioutil.WriteFile("examples/Test.png", b, 0644); err != nil {
		t.Error(err)
	}
}
