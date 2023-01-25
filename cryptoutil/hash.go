package cryptoutil

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/subtle"
	"encoding/base64"
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

func Encrypt(text, salt string, key *rsa.PublicKey, fn func() hash.Hash) (string, error) {
	cipher, err := rsa.EncryptOAEP(fn(), rand.Reader, key, []byte(text), []byte(salt))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}

func Decrypt(cipher, salt string, key *rsa.PrivateKey, fn func() hash.Hash) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", err
	}

	text, err := rsa.DecryptOAEP(fn(), rand.Reader, key, b, []byte(salt))
	if err != nil {
		return "", err
	}

	return string(text), err
}
