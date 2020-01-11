package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllArticles(t *testing.T) {
	MockArticles()
	defer UndoMockArticles()
	sOfArticles, err := GetAllArticles()

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
	MockEmptyArticles()
	defer UndoMockArticles()

	sOfArticles, err := GetAllArticles()

	assert.Nil(t, sOfArticles)
	assert.EqualError(t, err, "No articles found")
}

func TestGetArticle(t *testing.T) {
	MockArticles()
	defer UndoMockArticles()
	id := 1
	article, err := GetArticle(id)

	assert.Nil(t, err)
	assert.NotNil(t, article)

	assert.EqualValues(t, articles[1].ID, article.ID)
	assert.EqualValues(t, articles[1].Title, article.Title)
	assert.EqualValues(t, articles[1].Content, article.Content)
}

func TestNotFoundArticle(t *testing.T) {
	MockArticles()
	defer UndoMockArticles()
	id := 3
	article, err := GetArticle(id)

	assert.Nil(t, article)
	assert.EqualError(t, err, fmt.Sprintf("Article with id = %d not found", id))
}
