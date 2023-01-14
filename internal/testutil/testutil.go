package testutil

import (
	"encoding/hex"
	"math/rand"
	"testing"
)

func RandomString(t *testing.T, prefix string, n int) string {
	t.Helper()

	b := make([]byte, n/2, n/2)
	_, _ = rand.Read(b)
	return prefix + hex.EncodeToString(b)
}
