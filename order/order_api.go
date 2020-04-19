package order

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// OrderAPI ...
type OrderAPI struct {
	OrderSerive OrderSerive
}

// ProvideOrderAPI ...
func ProvideOrderAPI(o OrderSerive) OrderAPI {
	return OrderAPI{OrderSerive: o}
}

// GetOrdersByUserID ...
func (o *OrderAPI) GetOrdersByUserID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	o.OrderSerive.GetOrdersByUserID(userID)
}
