package file

import (
	"fmt"
	"image"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/llgcode/draw2d/draw2dsvg"
	"github.com/lobre/rm"
)

type Drawer struct {
	Notebook      *rm.Notebook
	TemplatesPath string
}

func NewDrawer(notebook *rm.Notebook) Drawer {
	return Drawer{notebook, ""}
}

func (d Drawer) DrawPng() error {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, rm.Width, rm.Height))
	gc := draw2dimg.NewGraphicContext(dest)

	// Draw content
	err := d.draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.png", d.Notebook.Name)
	err = draw2dimg.SaveToPngFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func (d Drawer) DrawPdf() error {
	// Initialize the graphic context on a pdf document
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Draw content
	err := d.draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.pdf", d.Notebook.Name)
	err = draw2dpdf.SaveToPdfFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func (d Drawer) DrawSvg() error {
	// Initialize the graphic context on an pdf document
	dest := draw2dsvg.NewSvg()
	gc := draw2dsvg.NewGraphicContext(dest)

	// Draw content
	err := d.draw(gc)
	if err != nil {
		return err
	}

	// Save to png
	fn := fmt.Sprintf("%s.svg", d.Notebook.Name)
	err = draw2dsvg.SaveToSvgFile(fn, dest)
	if err != nil {
		return err
	}

	return nil
}

func (d Drawer) draw(gc draw2d.GraphicContext) error {
	return nil
}
