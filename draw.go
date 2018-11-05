package notebook

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/llgcode/draw2d/draw2dsvg"
)

func (n *Notebook) DrawPng() ([]byte, error) {
	var f []byte

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, Width, Height))
	gc := draw2dimg.NewGraphicContext(dest)

	// Draw content
	err := n.draw(gc)
	if err != nil {
		return f, err
	}

	// Generate tmp file
	fn := fmt.Sprintf("%s.png", n.Name)
	tmp, err := ioutil.TempFile("", fn)
	if err != nil {
		return f, err
	}
	defer os.Remove(tmp.Name())

	// Save to png
	err = draw2dimg.SaveToPngFile(tmp.Name(), dest)
	if err != nil {
		return f, err
	}

	// Get bytes from tmp file
	f, err = ioutil.ReadFile(tmp.Name())
	if err != nil {
		return f, err
	}

	return f, nil
}

func (n *Notebook) DrawPdf() ([]byte, error) {
	var f []byte

	// Initialize the graphic context on a pdf document
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Draw content
	err := n.draw(gc)
	if err != nil {
		return f, err
	}

	// Generate tmp file
	fn := fmt.Sprintf("%s.pdf", n.Name)
	tmp, err := ioutil.TempFile("", fn)
	if err != nil {
		return f, err
	}
	defer os.Remove(tmp.Name())

	// Save to pdf
	err = draw2dpdf.SaveToPdfFile(tmp.Name(), dest)
	if err != nil {
		return f, err
	}

	// Get bytes from tmp file
	f, err = ioutil.ReadFile(tmp.Name())
	if err != nil {
		return f, err
	}

	return f, nil
}

func (n *Notebook) DrawSvg() ([]byte, error) {
	var f []byte

	// Initialize the graphic context on an pdf document
	dest := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(dest)

	// Draw content
	err := n.draw(gc)
	if err != nil {
		return f, err
	}

	// Generate tmp file
	fn := fmt.Sprintf("%s.svg", n.Name)
	tmp, err := ioutil.TempFile("", fn)
	if err != nil {
		return f, err
	}
	defer os.Remove(tmp.Name())

	// Save to png
	err = draw2dsvg.SaveToSvgFile(tmp.Name(), dest)
	if err != nil {
		return f, err
	}

	// Get bytes from tmp file
	f, err = ioutil.ReadFile(tmp.Name())
	if err != nil {
		return f, err
	}

	return f, nil
}

func (n *Notebook) draw(gc draw2d.GraphicContext) error {
	return nil
}
