package rm

import (
	"os"
)

// Width and Height of the device
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
	layers []layer
	//template  string
	//thumbnail image.Image
}

type contentTransform struct {
	M11 int `json:"m11"`
	M12 int `json:"m12"`
	M13 int `json:"m13"`
	M21 int `json:"m21"`
	M22 int `json:"m22"`
	M23 int `json:"m23"`
	M31 int `json:"m31"`
	M32 int `json:"m32"`
	M33 int `json:"m33"`
}

type contentExtraMetadata struct {
	LastColor      string `json:"LastColor"`
	LastTool       string `json:"LastTool"`
	ThicknessScale string `json:"ThicknessScale"`
}

type content struct {
	ExtraMetadata  contentExtraMetadata `json:"extraMetadata"`
	FileType       string               `json:"fileType"`
	FontName       string               `json:"fontName"`
	LastOpenedPage int                  `json:"lastOpenedPage"`
	LineHeight     int                  `json:"lineHeight"`
	Margins        int                  `json:"margins"`
	PageCount      int                  `json:"pageCount"`
	TextScale      int                  `json:"textScale"`
	Transform      contentTransform     `json:"transform"`
}

// Notebook parsed from the reMarkable
type Notebook struct {
	Name string

	id      string
	pages   []page
	content content
	pdf     os.File
	epub    os.File
	hash    string
}

const header = "reMarkable lines with selections and layers"
