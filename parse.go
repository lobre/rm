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
	if _, err := r.Read(b); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func readFloat32(r io.Reader) (float32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits), nil
}

func (n *Notebook) parseLines(r io.Reader) error {
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

	var pages []page

	for pidx := uint32(1); pidx <= nbPages; pidx++ {

		// Get number of layers
		nbLayers, err := readInt32(r)
		if err != nil {
			return err
		}

		var layers []layer

		for lyidx := uint32(1); lyidx <= nbLayers; lyidx++ {
			// Get number of lines
			nbLines, err := readInt32(r)
			if err != nil {
				return err
			}

			var lines []line

			for lidx := uint32(1); lidx <= nbLines; lidx++ {
				brshType, err := readInt32(r)
				if err != nil {
					return err
				}

				brshColor, err := readInt32(r)
				if err != nil {
					return err
				}

				_, err = readInt32(r) // Select and transform tool not used so far
				if err != nil {
					return err
				}

				brshSize, err := readFloat32(r)
				if err != nil {
					return err
				}

				// Get number of points
				nbPoints, err := readInt32(r)
				if err != nil {
					return err
				}

				var points []point

				for ptidx := uint32(1); ptidx <= nbPoints; ptidx++ {
					x, err := readFloat32(r)
					if err != nil {
						return err
					}

					y, err := readFloat32(r)
					if err != nil {
						return err
					}

					penPressure, err := readFloat32(r)
					if err != nil {
						return err
					}

					xRotation, err := readFloat32(r)
					if err != nil {
						return err
					}

					yRotation, err := readFloat32(r)
					if (err == io.EOF && ptidx < nbPoints) || (err != nil && err != io.EOF) {
						return err
					}

					points = append(points, point{x, y, penPressure, xRotation, yRotation})
				}

				lines = append(lines, line{
					brushType(brshType),
					brushColor(brshColor),
					brushSize(brshSize),
					points,
				})

			}

			layers = append(layers, layer{lines})
		}

		pages = append(pages, page{layers})
	}

	n.pages = pages
	return nil
}

func NewNotebook(zipFile string) (*Notebook, error) {
	name := filepath.Base(strings.TrimSuffix(zipFile, filepath.Ext(zipFile)))
	notebook := Notebook{Name: name}

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, err
	}

	// Calculate zip hash
	fr, err := os.Open(zipFile)
	h, err := md5Hash(fr)
	if err != nil {
		return nil, err
	}
	notebook.hash = h

	// Search for lines file
	for _, zf := range r.File {
		name := zf.FileInfo().Name()
		ext := filepath.Ext(name)

		if ext == ".lines" {
			f, err := zf.Open()
			if err != nil {
				return nil, err
			}
			defer f.Close()
			if err := notebook.parseLines(f); err != nil {
				return nil, fmt.Errorf("Can't parse lines file: %v", err)
			}

			// Set id
			notebook.id = strings.TrimSuffix(name, ext)
			break
		}
	}
	if len(notebook.pages) == 0 {
		return nil, errors.New("Notebook does not contain data")
	}

	// Process other files
	for _, zf := range r.File {
		name := zf.FileInfo().Name()
		ext := filepath.Ext(name)

		if ext == ".content" {
			f, err := zf.Open()
			if err != nil {
				return nil, err
			}
			defer f.Close()
			b, err := ioutil.ReadAll(f)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(b, &notebook.content)
		}
	}

	// Save thumbnails
	// for _, f := range r.File {
	// }

	return &notebook, nil
}
