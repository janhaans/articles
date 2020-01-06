package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tmpArticles Articles
)

func TestGetAllArticles(t *testing.T) {
	SaveArticles()
	MockArticles()
	defer RestoreArticles()

	sOfArticles, err := GetAllArticles()

	//fmt.Println(articles)
	//fmt.Println(sOfArticles)

	assert.Nil(t, err)
	assert.NotNil(t, sOfArticles)
	assert.EqualValues(t, len(articles), len(sOfArticles))
	for k, v := range articles {
		assert.EqualValues(t, v.ID, sOfArticles[k-1].ID)
		assert.EqualValues(t, v.Title, sOfArticles[k-1].Title)
		assert.EqualValues(t, v.Content, sOfArticles[k-1].Content)
	}
}
func TestEmptyArticles(t *testing.T) {
	SaveArticles()
	MockEmptyArticles()
	defer RestoreArticles()

	sOfArticles, err := GetAllArticles()

	assert.Nil(t, sOfArticles)
	assert.EqualError(t, err, "No articles found")
}

func TestGetArticle(t *testing.T) {
	SaveArticles()
	MockArticles()
	defer RestoreArticles()

	id := 1
	article, err := GetArticle(id)

	assert.Nil(t, err)
	assert.NotNil(t, article)

	assert.EqualValues(t, articles[1].ID, article.ID)
	assert.EqualValues(t, articles[1].Title, article.Title)
	assert.EqualValues(t, articles[1].Content, article.Content)
}

func TestNotFoundArticle(t *testing.T) {
	SaveArticles()
	MockArticles()
	defer RestoreArticles()

	id := 3
	article, err := GetArticle(id)

	assert.Nil(t, article)
	assert.EqualError(t, err, fmt.Sprintf("No article with id = %d found", id))
}

func SaveArticles() {
	tmpArticles = articles
}

func RestoreArticles() {
	articles = tmpArticles
}

func MockEmptyArticles() {
	articles = Articles{}
}

func MockArticles() {
	articles = Articles{
		1: &Article{1, "mock title 1", "mock description 1"},
		2: &Article{2, "mock title 2", "mock description 2"},
	}
}
