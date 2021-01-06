package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatal("Failed : ", res.Code)
	}

	assert.Equal(http.StatusOK, res.Code)
}
