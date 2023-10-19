package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

func TestAs(t *testing.T) {
	router := RegisterRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/transaction", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}
