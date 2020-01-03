package app

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	mapURL(router)
	return router
}

//StartApp starts the articles application
func StartApp() {
	router := setupRouter()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
