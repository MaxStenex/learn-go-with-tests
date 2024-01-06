package context

import (
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			w.Write([]byte(d))
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
