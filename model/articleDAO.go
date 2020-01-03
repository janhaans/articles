package model

import (
	"errors"
)

//Articles simulate database
type Articles map[int]Article

var articles = Articles{
	1: Article{1, "title 1", "description 1"},
	2: Article{2, "title 2", "description 2"},
}

//GetAllArticles provides all articles
func GetAllArticles() (Articles, error) {
	if len(articles) == 0 {
		return nil, errors.New("No articles found")
	}
	return articles, nil
}
