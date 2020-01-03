package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janhaans/articles/model"
)

//ShowIndexPage is the index page and shows all articles
func ShowIndexPage(c *gin.Context) {
	articles, _ := model.GetAllArticles()
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": articles})
}
