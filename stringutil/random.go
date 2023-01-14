package stringutil

import (
	"crypto/rand"
	"encoding/hex"
)

func Random(n int) string {
	if n <= 0 {
		return ""
	}

	b := make([]byte, n/2, n/2)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
