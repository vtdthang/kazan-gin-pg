package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vtdthang/kazan-gin-pg/order"
)

// PrepareOrderRoutes ...
func PrepareOrderRoutes(route *gin.Engine, o order.OrderAPI) {
	productRoute := route.Group("/api/v1/users")
	productRoute.GET("/:user_id/orders", o.GetOrdersByUserID)
}
