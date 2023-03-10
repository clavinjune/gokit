package cryptoutil_test

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"

	"github.com/clavinjune/gokit/cryptoutil"
	"github.com/stretchr/testify/require"
)

func TestHashVerify(t *testing.T) {
	tt := []struct {
		_      struct{}
		Name   string
		Text   string
		Salt   string
		HashFn func() hash.Hash
	}{
		{
			Name:   "SHA256 without salt",
			Text:   "foobar",
			Salt:   "",
			HashFn: sha256.New,
		},
		{
			Name:   "SHA256 with salt",
			Text:   "foobar",
			Salt:   "salty",
			HashFn: sha256.New,
		},
		{
			Name:   "MD5 without salt",
			Text:   "foobar",
			Salt:   "",
			HashFn: md5.New,
		},
		{
			Name:   "MD5 with salt",
			Text:   "foobar",
			Salt:   "salty",
			HashFn: md5.New,
		},
		{
			Name:   "SHA512 without salt",
			Text:   "foobar",
			Salt:   "",
			HashFn: sha512.New,
		},
		{
			Name:   "SHA512 with salt",
			Text:   "foobar",
			Salt:   "salty",
			HashFn: sha512.New,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)
			cipher := cryptoutil.HashGenerate(tc.Text, tc.Salt, tc.HashFn)
			r.True(cryptoutil.HashVerify(tc.Text, cipher, tc.Salt, tc.HashFn))
		})
	}
}
