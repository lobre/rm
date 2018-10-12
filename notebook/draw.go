package notebook

import (
	"fmt"
	"image"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/llgcode/draw2d/draw2dsvg"
)

func (n *Notebook) DrawPng() error {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, Width, Height))
	gc := draw2dimg.NewGraphicContext(dest)

	// Draw content
	err := draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.png", n.Name)
	err = draw2dimg.SaveToPngFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notebook) DrawPdf() error {
	// Initialize the graphic context on a pdf document
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Draw content
	err := draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.pdf", n.Name)
	err = draw2dpdf.SaveToPdfFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notebook) DrawSvg() error {
	// Initialize the graphic context on an pdf document
	dest := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(dest)

	// Draw content
	err := draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.svg", n.Name)
	err = draw2dsvg.SaveToSvgFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func draw(gc draw2d.GraphicContext) error {
	return nil
}
