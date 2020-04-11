package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vtdthang/kazan-gin-pg/product"
)

// PrepareProductRoutes ...
func PrepareProductRoutes(route *gin.Engine, p product.ProductAPI) {
	productRoute := route.Group("/api/v1/products")
	productRoute.GET("/", p.FindAll)
	productRoute.GET("/:id", p.FindByID)
	productRoute.POST("/", p.Create)
	productRoute.PUT("/:id", p.Update)
	productRoute.DELETE("/:id", p.Delete)
}
