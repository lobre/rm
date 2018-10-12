package notebook

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func md5Hash(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	hb := h.Sum(nil)[:16]
	hs := hex.EncodeToString(hb)
	return hs, nil
}
