package server

import (
  "github.com/gin-gonic/gin"
)

// Router .
func Router() *gin.Engine  {

  router := gin.New()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())

  // System routes
  routes.NotFound(router)
  routes.NoMethods(router)

  return router
}