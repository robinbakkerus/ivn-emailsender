package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router .
func Router() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// System routes
	NotFound(router)
	NoMethods(router)

	router.LoadHTMLGlob("./web/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return router
}
