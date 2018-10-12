package notebook

import (
	"strings"
	"testing"
)

// The hash has been generated using
// echo -n "test" | md5sum
func TestMd5Hash(t *testing.T) {
	r := strings.NewReader("test")
	h := "098f6bcd4621d373cade4e832627b4f6"
	hash, err := md5Hash(r)
	if err != nil {
		t.Errorf("Impossible to calculate hash: %v", err)
	}
	if hash != h {
		t.Errorf("The hash is not correct")
	}
}
