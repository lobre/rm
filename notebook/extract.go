package notebook

import (
	"fmt"
	"io/ioutil"
)

func (n *Notebook) ExtractPdf(path string) error {
	if err := ioutil.WriteFile(path, n.pdf, 0644); err != nil {
		return err
	}
	return nil
}

func (n *Notebook) ExtractEpub(path string) error {
	if err := ioutil.WriteFile(path, n.epub, 0644); err != nil {
		return err
	}
	return nil
}

func (n *Notebook) ExtractThumbnail(i int, path string) error {
	if i >= len(n.Pages) {
		return fmt.Errorf("Page does not exist")
	}
	if n.Pages[i].thumbnail == nil {
		return fmt.Errorf("No thumbnail for this page")
	}
	if err := ioutil.WriteFile(path, n.Pages[i].thumbnail, 0644); err != nil {
		return err
	}
	return nil
}
