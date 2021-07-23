package _1_select_partone

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func OriginalTestRacer(t *testing.T) {

	slowServer := OriginalmakeDelayedServer(20 * time.Millisecond)
	fastServer := OriginalmakeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := OriginalRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func OriginalmakeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
