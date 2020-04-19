package entities

// OrderDetail ...
type OrderDetail struct {
	OrderID   int `gorm:"primary_key;auto_increment:false" json:"order_id"`
	Order     Order
	ProductID int `gorm:"primary_key;auto_increment:false" json:"product_id"`
	Product   Product
	Quantity  int `gorm:"not null" json:"quantity"`
}
