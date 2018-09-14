package main

import (
	"encoding/binary"
	"flag"
	"fmt"
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

func NewPoint(x, y, p, xr, yr float32) Point {
	return Point{x, y, p, xr, yr}
}

func NewLine(t BrushType, c BrushColor, s BrushSize, p []Point) Line {
	return Line{t, c, s, p}
}

func NewLayer(l []Line) Layer {
	return Layer{l}
}

func NewPage(l []Layer) Page {
	return Page{l}
}

func NewNotebook(n string, p []Page) Notebook {
	return Notebook{n, p}
}

func ReadInt32(f *os.File) int {
	b := make([]byte, 4)
	f.Read(b)
	return int(binary.LittleEndian.Uint32(b))
}

func ReadFloat32(f *os.File) float32 {
	b := make([]byte, 4)
	f.Read(b)
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits)
}

const Header = "reMarkable lines with selections and layers"

func parse(fileName string) Notebook {
	f, _ := os.Open(fileName)

	// Check header
	h := make([]byte, 43)
	f.Read(h)
	if string(h) != Header {
		fmt.Println("Wrong input file type")
		os.Exit(1)
	}

	// Get number of pages
	nbPages := ReadInt32(f)

	var pages []Page

	for pageIdx := 1; pageIdx <= nbPages; pageIdx++ {

		// Get number of layers
		nbLayers := ReadInt32(f)

		var layers []Layer

		for layerIdx := 1; layerIdx <= nbLayers; layerIdx++ {
			// Get number of lines
			nbLines := ReadInt32(f)

			var lines []Line

			for lineIdx := 1; lineIdx <= nbLines; lineIdx++ {
				brushType := ReadInt32(f)
				brushColor := ReadInt32(f)
				_ = ReadInt32(f) // Select and transform tool not used so far
				brushSize := ReadFloat32(f)

				// Get number of points
				nbPoints := ReadInt32(f)

				var points []Point

				for pointIdx := 1; pointIdx <= nbPoints; pointIdx++ {
					x := ReadFloat32(f)
					y := ReadFloat32(f)
					penPressure := ReadFloat32(f)
					xRotation := ReadFloat32(f)
					yRotation := ReadFloat32(f)

					points = append(points, NewPoint(x, y, penPressure, xRotation, yRotation))
				}

				lines = append(lines, NewLine(
					BrushType(brushType),
					BrushColor(brushColor),
					BrushSize(brushSize),
					points,
				))

			}

			layers = append(layers, NewLayer(lines))
		}

		pages = append(pages, NewPage(layers))
	}

	return NewNotebook(fileName, pages)
}

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
