package model

import (
	"errors"
	"fmt"
)

//Articles simulate database
type Articles map[int]*Article

var articles = Articles{
	1: &Article{1, "title 1", "description 1"},
	2: &Article{2, "title 2", "description 2"},
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
	return nil, fmt.Errorf("No article with id = %d found", id)
}
