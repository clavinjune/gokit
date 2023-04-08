package proxyutil

import "github.com/clavinjune/gokit/addressutil"

var (
	_ addressutil.Addresser = (*Address)(nil)
)

type Address string

func (a Address) Address() string {
	return string(a)
}
