package racer

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/select"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	timeout := 20 * time.Second

	t.Run("check Racer func", func(t *testing.T) {
		slowServer := makeTestServer(http.StatusOK, 20*time.Millisecond)
		fastServer := makeTestServer(http.StatusOK, 10*time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		exp := fastUrl
		got, _ := Racer(slowUrl, fastUrl, timeout)

		AssertEquals(t, exp, got)
	})

	t.Run("check Racer func with very slow servers", func(t *testing.T) {
		slowServer := makeTestServer(http.StatusOK, 20*time.Second)
		slowerServer := makeTestServer(http.StatusOK, 25*time.Second)

		defer slowServer.Close()
		defer slowerServer.Close()

		_, err := Racer(slowServer.URL, slowerServer.URL, timeout)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeTestServer(status int, delay time.Duration) (server *httptest.Server) {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(status)
	}))
}
