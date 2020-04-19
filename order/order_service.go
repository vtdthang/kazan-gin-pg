package order

// OrderSerive ...
type OrderSerive struct {
	OrderRepository OrderRepository
}

// ProvideOrderService ...
func ProvideOrderService(o OrderRepository) OrderSerive {
	return OrderSerive{OrderRepository: o}
}

// GetOrdersByUserID ...
func (o OrderSerive) GetOrdersByUserID(userID int) {
	o.OrderRepository.GetOrdersByUserID(userID)
}
