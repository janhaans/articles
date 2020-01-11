package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janhaans/articles/model"
)

//ShowIndexPage is the index page and shows all articles
func ShowIndexPage(c *gin.Context) {
	articles, err := model.GetAllArticles()
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{"title": "Article", "errorMsg": err.Error()})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": articles})
	}
}

//ShowArticlePage shows an article
func ShowArticlePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"title": "Article", "errorMsg": "Bad request: latest path element must be integer"})
	} else {
		article, err := model.GetArticle(id)
		if article == nil {
			c.HTML(http.StatusOK, "error.html", gin.H{"title": "Article", "errorMsg": err.Error()})
		} else {
			c.HTML(http.StatusOK, "article.html", gin.H{"title": "Article", "payload": article})
		}
	}
}
