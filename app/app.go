package app

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/janhaans/articles/utils/projectpath"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob(filepath.Join(projectpath.Root, "templates/*"))
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
