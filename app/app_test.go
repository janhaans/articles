package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/janhaans/articles/model"

	"github.com/stretchr/testify/assert"
)

func TestShowIndexPage(t *testing.T) {
	model.MockArticles()
	defer model.UndoMockArticles()
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	res := GetResponse(req)

	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.True(t, CheckResponseBody(res, "title 1", "description 1", "title 2", "description 2"))
}

func TestNoArticlesFound(t *testing.T) {
	model.MockEmptyArticles()
	defer model.UndoMockArticles()

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	res := GetResponse(req)

	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.True(t, CheckResponseBody(res, "No articles found"))
}

func TestGetArticle(t *testing.T) {
	model.MockArticles()
	defer model.UndoMockArticles()
	req, err := http.NewRequest("GET", "/article/view/1", nil)
	assert.Nil(t, err)
	res := GetResponse(req)

	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.True(t, CheckResponseBody(res, "Article 1", "title 1", "description 1"))
}

func TestNoArticleFound(t *testing.T) {
	model.MockArticles()
	defer model.UndoMockArticles()
	req, err := http.NewRequest("GET", "/article/view/0", nil)
	assert.Nil(t, err)
	res := GetResponse(req)

	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.True(t, CheckResponseBody(res, "Article with id = 0 not found"))
}

func TestBadArticleRequest(t *testing.T) {
	model.MockArticles()
	defer model.UndoMockArticles()
	req, err := http.NewRequest("GET", "/article/view/hello", nil)
	assert.Nil(t, err)
	res := GetResponse(req)

	assert.EqualValues(t, http.StatusBadRequest, res.Code)
	assert.EqualValues(t, "text/html; charset=utf-8", res.HeaderMap.Get("Content-Type"))
	assert.True(t, CheckResponseBody(res, "Bad request: latest path element must be integer"))
}

func GetResponse(req *http.Request) *httptest.ResponseRecorder {
	router := setupRouter()
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func CheckResponseBody(res *httptest.ResponseRecorder, substrs ...string) bool {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}
	for _, substr := range substrs {
		if strings.Index(string(body), substr) < 0 {
			return false
		}
	}
	return true
}
