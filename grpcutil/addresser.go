package grpcutil

import "fmt"

type Addresser interface {
	Address() string
}

type ServerPort int

func (s ServerPort) Address() string {
	return fmt.Sprintf(":%d", s)
}

type ProxyPort int

func (p ProxyPort) Address() string {
	return fmt.Sprintf(":%d", p)
}
