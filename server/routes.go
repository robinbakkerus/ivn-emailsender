package server

import "github.com/gin-gonic/gin"

func NotFound(route *gin.Engine) {
  route.NoRoute(func(c *gin.Context) {
     c.AbortWithStatusJSON(404, "Not Found")
  })
}

func NoMethods(route *gin.Engine){
  route.NoMethod(func(c *gin.Context) {
     c.AbortWithStatusJSON(405, "not allowed")
  })
}