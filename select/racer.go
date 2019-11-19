package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) string {

	dur1 := measureTime(url1)
	dur2 := measureTime(url2)

	if dur1 < dur2 {
		return url1
	}
	return url2
}

func Racer2(url1 string, url2 string, timeout time.Duration) (string, error) {

	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errors.New("lofasz!")
	}
}

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
