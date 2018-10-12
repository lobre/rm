package notebook

import (
	"testing"
)

func TestDrawPng(t *testing.T) {
	// Parse notebook from zip file
	n, err := NewNotebook("../examples/Test.zip")
	if err != nil {
		t.Errorf("Impossible to create notebook from zip: %v", err)
	}

	// Draw png file
	err = n.DrawPng()
	if err != nil {
		t.Errorf("Error while generating png: %v", err)
	}
}
