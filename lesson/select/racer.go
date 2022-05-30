package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()

	return ch
}
