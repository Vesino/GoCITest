package apiserver

import "net/http"

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(("Hello gopher!")))
	}
}
