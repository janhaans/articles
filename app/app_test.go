package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowIndexPage(t *testing.T) {
	router := setupRouter()
	req, err := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
}

func TestShowArticlePage(t *testing.T) {
	router := setupRouter()
	req, err := http.NewRequest("GET", "/article/view/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
}
