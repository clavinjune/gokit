package proxyutil

import "github.com/clavinjune/gokit/etcutil"

var (
	_ etcutil.Addresser = (*Address)(nil)
)

type Address string

func (a Address) Address() string {
	return string(a)
}
