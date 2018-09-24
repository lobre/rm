package rm

import (
	"image"
	"os"
)

const (
	Width  int = 1404
	Height int = 1872
)

type brushColor int
type brushType int
type brushSize float32

// Brush colors
const (
	black brushColor = 0
	grey  brushColor = 1
	white brushColor = 2
)

// Brush types
const (
	brush       brushType = 0
	pencilTilt  brushType = 1
	pen         brushType = 2
	marker      brushType = 3
	fineliner   brushType = 4
	highlighter brushType = 5
	eraser      brushType = 6
	pencilSharp brushType = 7
	eraseArea   brushType = 8
)

// Brush sizes
const (
	small  brushSize = 1.875
	medium brushSize = 2.0
	large  brushSize = 2.125
)

type point struct {
	x           float32
	y           float32
	penPressure float32
	xRotation   float32
	yRotation   float32
}

type line struct {
	brushType  brushType
	brushColor brushColor
	brushSize  brushSize
	points     []point
}

type layer struct {
	lines []line
}

type page struct {
	layers    []layer
	template  string
	thumbnail image.Image
}

// Notebook parsed from the reMarkable
type Notebook struct {
	Name string

	id      string
	pages   []page
	content map[string]interface{}
	pdf     os.File
	epub    os.File
}

const header = "reMarkable lines with selections and layers"
