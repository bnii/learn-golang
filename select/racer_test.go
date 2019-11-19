package racer

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowserver := makeDelayedServer(20 * time.Millisecond)
	fastserver := makeDelayedServer(0 * time.Millisecond)

	defer slowserver.Close()
	defer fastserver.Close()

	slowURL := slowserver.URL
	fastURL := fastserver.URL
	got := Racer(fastURL, slowURL)
	assert.Equal(t, fastURL, got)

}

func TestRaceChan(t *testing.T) {

	t.Run("std cchannel racer", func(t *testing.T) {
		slowserver := makeDelayedServer(20 * time.Millisecond)
		fastserver := makeDelayedServer(0 * time.Millisecond)

		defer slowserver.Close()
		defer fastserver.Close()

		slowURL := slowserver.URL
		fastURL := fastserver.URL
		got, _ := Racer2(fastURL, slowURL, 1*time.Second)
		assert.Equal(t, fastURL, got)
	})

	t.Run("returns error for more than 10 sec", func(t *testing.T) {
		slowserver := makeDelayedServer(12 * time.Second)
		fastserver := makeDelayedServer(11 * time.Second)

		defer slowserver.Close()
		defer fastserver.Close()

		slowURL := slowserver.URL
		fastURL := fastserver.URL
		_, err := Racer2(fastURL, slowURL, 20*time.Millisecond)
		assert.Error(t, err, "timeout exceeded")
	})

}

func makeDelayedServer(i time.Duration) *httptest.Server {
	slowserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(i)
		w.WriteHeader(http.StatusOK)
	}))
	return slowserver
}
