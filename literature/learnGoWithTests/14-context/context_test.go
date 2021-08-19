package _4_context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

// A store interface defines anything with a Fetch func which returns a string
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("tell the store to cancel work if the request is cancelled", func(*testing.T) {
		data := "Hello, World"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was told not to cancel")
		}

	})

	t.Run("returns data from the store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		if !store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})

}
