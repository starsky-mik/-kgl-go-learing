package context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		_, _ = fmt.Fprint(w, res)
	}
}
