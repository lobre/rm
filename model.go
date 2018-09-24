package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
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
	X           float32
	Y           float32
	PenPressure float32
	XRotation   float32
	YRotation   float32
}

type Line struct {
	BrushType  BrushType
	BrushColor BrushColor
	BrushSize  BrushSize
	Points     []Point
}

type Layer struct {
	Lines []Line
}

type Page struct {
	Layers []Layer
}

type Notebook struct {
	Name  string
	Pages []Page
}

func readInt32(r io.Reader) uint32 {
	b := make([]byte, 4)
	r.Read(b)
	return binary.LittleEndian.Uint32(b)
}

func readFloat32(r io.Reader) float32 {
	b := make([]byte, 4)
	r.Read(b)
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits)
}

const header = "reMarkable lines with selections and layers"

func NewNotebook(r io.Reader, name string) (*Notebook, error) {
	// Check header
	h := make([]byte, len(header))
	r.Read(h)
	if string(h) != header {
		fmt.Println("Wrong input file type")
		os.Exit(1)
	}

	// Get number of pages
	nbPages := readInt32(r)

	var pages []Page

	for pidx := uint32(1); pidx <= nbPages; pidx++ {

		// Get number of layers
		nbLayers := readInt32(r)

		var layers []Layer

		for lyidx := uint32(1); lyidx <= nbLayers; lyidx++ {
			// Get number of lines
			nbLines := readInt32(r)

			var lines []Line

			for lidx := uint32(1); lidx <= nbLines; lidx++ {
				brushType := readInt32(r)
				brushColor := readInt32(r)
				_ = readInt32(r) // Select and transform tool not used so far
				brushSize := readFloat32(r)

				// Get number of points
				nbPoints := readInt32(r)

				var points []Point

				for ptidx := uint32(1); ptidx <= nbPoints; ptidx++ {
					x := readFloat32(r)
					y := readFloat32(r)
					penPressure := readFloat32(r)
					xRotation := readFloat32(r)
					yRotation := readFloat32(r)

					points = append(points, Point{x, y, penPressure, xRotation, yRotation})
				}

				lines = append(lines, Line{
					BrushType(brushType),
					BrushColor(brushColor),
					BrushSize(brushSize),
					points,
				})

			}

			layers = append(layers, Layer{lines})
		}

		pages = append(pages, Page{layers})
	}

	return &Notebook{name, pages}, nil
}
