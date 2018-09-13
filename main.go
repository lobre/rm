package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

type BrushColor int
type BrushType int
type BrushSize float32

const (
	Black BrushColor = 0
	Grey  BrushColor = 1
	White BrushColor = 2
)

const (
	Brush       BrushType = 0
	PencilTilt  BrushType = 1
	Pen         BrushType = 2
	Marker      BrushType = 3
	Fineliner   BrushType = 4
	Highlighter BrushType = 5
	Eraser      BrushType = 6
	PencilSharp BrushType = 7
	EraseArea   BrushType = 8
)

const (
	Small  BrushSize = 1.875
	Medium BrushSize = 2.0
	Large  BrushSize = 2.125
)

type Point struct {
}

type Line struct {
	BrushType  BrushType
	BrushColor BrushColor
}

type Layer struct {
}

type Page struct {
}

type Notebook struct {
	Layers []Layer
}

func NewLine(b BrushType, c BrushColor) Line {
	return Line{b, c}
}

func NewNotebook() Notebook {
	return Notebook{}
}

func ReadInt32(f *os.File) int {
	b := make([]byte, 4)
	f.Read(b)
	return int(binary.LittleEndian.Uint32(b))
}

const Header = "reMarkable lines with selections and layers"

func main() {
	// Get filename from flags
	var file = flag.String("file", "", "provide a lines file")
	flag.Parse()
	if *file == "" {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	f, _ := os.Open(*file)

	// Check header
	h := make([]byte, 43)
	f.Read(h)
	if string(h) != Header {
		fmt.Println("Wrong input file type")
		os.Exit(1)
	}

	// Initialize new Notebook
	notebook := NewNotebook()

	// Get number of pages
	nbPages := ReadInt32(f)

	for pageIdx := 1; pageIdx <= nbPages; pageIdx++ {
		// Get number of layers
		nbLayers := ReadInt32(f)

		for layerIdx := 1; layerIdx <= nbLayers; layerIdx++ {
			// Get number of lines
			nbLines := ReadInt32(f)

			for lineIdx := 1; lineIdx <= nbLines; lineIdx++ {
				brushType := ReadInt32(f)
				brushColor := ReadInt32(f)
				// Select and transform tool not used so far
				_ = ReadInt32(f)

				line := NewLine(
					BrushType(brushType),
					BrushColor(brushColor),
				)

				fmt.Printf("%#v", line)
			}

		}
	}

	fmt.Printf("%#v", notebook)
}
