package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janhaans/articles/model"
)

//ShowIndexPage is the index page and shows all articles
func ShowIndexPage(c *gin.Context) {
	articles, _ := model.GetAllArticles()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": articles})
}

func ShowArticlePage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, _ := model.GetArticle(id)
	c.HTML(http.StatusOK, "article.html", gin.H{"title": "Article", "payload": article})
}
