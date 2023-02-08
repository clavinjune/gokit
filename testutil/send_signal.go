package testutil

import (
	"os"
	"testing"
)

func SendSignal(t *testing.T, ch <-chan struct{}, s os.Signal) {
	t.Helper()

	<-ch
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err.Error())
	}

	if err := p.Signal(s); err != nil {
		t.Fatal(err.Error())
	}
}
