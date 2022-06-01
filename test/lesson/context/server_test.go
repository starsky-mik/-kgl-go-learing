package context

import (
	"context"
	"errors"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	res string
	t   *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	ch := make(chan string, 1)

	go func() {
		var res string

		for _, c := range s.res {
			select {
			case <-ctx.Done():
				log.Println("canceled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				res += string(c)
			}
		}

		ch <- res
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("return data", func(t *testing.T) {
		want := "Hello!"
		store := SpyStore{res: want}
		srv := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		srv.ServeHTTP(res, req)

		got := res.Body.String()

		AssertEquals(t, got, want)
	})

	t.Run("cancel request", func(t *testing.T) {
		want := "Hello!"
		store := SpyStore{res: want}
		srv := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cCtx)

		res := &SpyResponseWriter{}

		srv.ServeHTTP(res, req)

		if res.written {
			t.Error("response written")
		}
	})
}
