package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"hash"
)

func GenerateKey(bits int) (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errBase, err)
	}

	if err := key.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", errBase, err)
	}

	return key, nil
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
