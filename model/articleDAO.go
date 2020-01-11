package model

import (
	"errors"
	"fmt"
)

//Articles simulate database
type Articles map[int]*Article

var (
	articles = Articles{
		1: &Article{1, "title 1", "description 1"},
		2: &Article{2, "title 2", "description 2"},
	}
	tmpArticles Articles
)

//SaveArticles saves articles before mocking articles
func SaveArticles() {
	tmpArticles = articles
}

//RestoreArticles restores articles after mocking articles
func RestoreArticles() {
	articles = tmpArticles
}

//MockEmptyArticles mocks an empty articles
func MockEmptyArticles() {
	SaveArticles()
	articles = Articles{}
}

//MockArticles mocks articles
func MockArticles() {
	SaveArticles()
	articles = Articles{
		1: &Article{1, "title 1", "description 1"},
		2: &Article{2, "title 2", "description 2"},
	}
}

//UndoMockArticles restores articles
func UndoMockArticles() {
	RestoreArticles()
}

//GetAllArticles provides all articles
func GetAllArticles() ([]*Article, error) {
	if len(articles) == 0 {
		return nil, errors.New("No articles found")
	}
	sOfArticles := make([]*Article, 0)
	for _, v := range articles {
		sOfArticles = append(sOfArticles, v)
	}
	return sOfArticles, nil
}

//GetArticle return article with requested id
func GetArticle(id int) (*Article, error) {
	article, ok := articles[id]
	if ok {
		return article, nil
	}
	return nil, fmt.Errorf("Article with id = %d not found", id)
}
