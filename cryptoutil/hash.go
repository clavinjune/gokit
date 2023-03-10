package cryptoutil

import (
	"crypto/hmac"
	"crypto/subtle"
	"encoding/hex"
	"hash"
)

func HashGenerate(text, salt string, fn func() hash.Hash) string {
	mac := hmac.New(fn, []byte(salt))
	mac.Write([]byte(text))

	return hex.EncodeToString(mac.Sum(nil))
}

func HashVerify(text, cipherText, salt string, fn func() hash.Hash) bool {
	actual := HashGenerate(text, salt, fn)
	return subtle.ConstantTimeCompare([]byte(actual), []byte(cipherText)) == 1
}
