package util

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

// Check will panic if the given error isn't nil.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// ExecuteRequest sends the given request to the given router and return the
// response.
func ExecuteRequest(req *http.Request, r *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}
