package router

import "github.com/gin-gonic/gin"

// PrepareUserRoutes ...
func PrepareUserRoutes(route *gin.Engine) {
	userRoute := route.Group("/api/v1/users")

	userRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
}
