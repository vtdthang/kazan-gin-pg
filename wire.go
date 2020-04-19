//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/vtdthang/kazan-gin-pg/order"
	"github.com/vtdthang/kazan-gin-pg/product"
)

// InitProductAPI ...
func InitProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}

func InitOrderAPI(db *gorm.DB) order.OrderAPI {
	wire.Build(order.ProvideOrderRepository, order.ProvideOrderService, order.ProvideOrderAPI)

	return order.OrderAPI{}
}
