package rm

import (
	"archive/zip"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func readInt32(r io.Reader) (uint32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil && err != io.EOF {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func readFloat32(r io.Reader) (float32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil && err != io.EOF {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits), nil
}

func (n *Notebook) parseLines(r io.Reader) error {
	log.Debug("START PARSING LINES FILE")
	log.Debug("--------------")

	// Check header
	h := make([]byte, len(header))
	if _, err := r.Read(h); err != nil {
		return err
	}
	if string(h) != header {
		return errors.New("Wrong input file type")
	}

	// Get number of pages
	nbPages, err := readInt32(r)

	if err != nil {
		return err
	}

	var pages []Page

	for pidx := uint32(1); pidx <= nbPages; pidx++ {
		log.Debugf("Page [%d/%d]", pidx, nbPages)

		// Get number of layers
		nbLayers, err := readInt32(r)
		if err != nil {
			return err
		}

		var layers []Layer

		for lyidx := uint32(1); lyidx <= nbLayers; lyidx++ {
			log.Debugf(indent(1, "Layer [%d/%d]"), lyidx, nbLayers)

			// Get number of lines
			nbLines, err := readInt32(r)
			if err != nil {
				return err
			}

			var lines []Line

			for lidx := uint32(1); lidx <= nbLines; lidx++ {
				log.Debugf(indent(2, "Line [%d/%d]"), lidx, nbLines)

				// Brush type
				bt, err := readInt32(r)
				if err != nil {
					return err
				}
				log.Debugf(indent(3, "Brush Type %d"), bt)

				// Brush color
				bc, err := readInt32(r)
				if err != nil {
					return err
				}
				log.Debugf(indent(3, "Brush Color %d"), bc)

				_, err = readInt32(r) // Select and transform tool not used so far
				if err != nil {
					return err
				}

				// Brush size
				bs, err := readFloat32(r)
				if err != nil {
					return err
				}
				log.Debugf(indent(3, "Brush Size %f"), bs)

				// Get number of points
				nbPoints, err := readInt32(r)
				if err != nil {
					return err
				}

				var points []Point

				for ptidx := uint32(1); ptidx <= nbPoints; ptidx++ {
					log.Debugf(indent(3, "Point [%d/%d]"), ptidx, nbPoints)

					x, err := readFloat32(r)
					if err != nil {
						return err
					}

					y, err := readFloat32(r)
					if err != nil {
						return err
					}
					log.Debugf(indent(4, "X, Y %f, %f"), x, y)

					penPressure, err := readFloat32(r)
					if err != nil {
						return err
					}
					log.Debugf(indent(4, "Pen Pressure %f"), y)

					xRotation, err := readFloat32(r)
					if err != nil {
						return err
					}
					log.Debugf(indent(4, "X Rotation %f"), xRotation)

					yRotation, err := readFloat32(r)
					if err != nil {
						return err
					}
					log.Debugf(indent(4, "Y Rotation %f"), yRotation)

					points = append(points, Point{x, y, penPressure, xRotation, yRotation})
				}

				lines = append(lines, Line{
					BrushType(bt),
					BrushColor(bc),
					BrushSize(bs),
					points,
				})

			}

			layers = append(layers, Layer{lines})
		}

		pages = append(pages, Page{layers, "", nil})
	}

	n.Pages = pages
	log.Debug("--------------")
	log.Debug("END OF PARSING")
	return nil
}

func (n *Notebook) parseContent(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &n.Content)
	return nil
}

// TODO
func (n *Notebook) parsePagedata(r io.Reader) error {
	return nil
}

// TODO
func (n *Notebook) parsePdf(r io.Reader) error {
	return nil
}

// TODO
func (n *Notebook) parseEpub(r io.Reader) error {
	return nil
}

// TODO
func (n *Notebook) parseThumbnails(zf []*zip.File) error {
	return nil
}

func NewNotebook(zipFile string) (*Notebook, error) {
	notebook := Notebook{}

	// Get name from zip filename
	notebook.Name = fileRawName(zipFile)

	// Calculate zip hash
	fr, err := os.Open(zipFile)
	h, err := md5Hash(fr)
	if err != nil {
		return nil, err
	}
	notebook.Hash = h

	// Get zip file reader
	zr, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}

	// Make sure zip contains files and register uuid
	zf := zipFirstFile(zr.File)
	if zf == nil {
		return nil, fmt.Errorf("No file in zip")
	}
	notebook.UUID = fileRawName(zf.FileInfo().Name())

	// Process lines file
	zf = zipSearchExt(".lines", zr.File)
	if zf != nil {
		f, err := zf.Open()
		if err != nil {
			return &notebook, err
		}
		defer f.Close()

		if err := notebook.parseLines(f); err != nil {
			return &notebook, fmt.Errorf("Can't parse lines file: %v", err)
		}
	}

	// Process content file
	zf = zipSearchExt(".content", zr.File)
	if zf != nil {
		f, err := zf.Open()
		if err != nil {
			return &notebook, err
		}
		defer f.Close()

		if err := notebook.parseContent(f); err != nil {
			return &notebook, fmt.Errorf("Can't parse content file: %v", err)
		}
	}

	return &notebook, nil
}

// Search a particular file with given extension in a zip
func zipSearchExt(ext string, files []*zip.File) *zip.File {
	// Search for lines file
	for _, zf := range files {
		name := zf.FileInfo().Name()
		e := filepath.Ext(name)
		if e == ext {
			return zf
		}
	}
	return nil
}

// Get first file of zip
func zipFirstFile(files []*zip.File) *zip.File {
	// Search for lines file
	for _, zf := range files {
		return zf
	}
	return nil
}

// Get filename without base nor extension
func fileRawName(s string) string {
	ext := filepath.Ext(s)
	return filepath.Base(strings.TrimSuffix(s, ext))
}
