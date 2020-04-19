package entities

// Product ...
type Product struct {
	ID           int           `gorm:"primary_key" json:"id"`
	CategoryID   int           `gorm:"index:idx_products_category_id" json:"category_id"`
	Category     Category      `gorm:"foreignkey:CategoryID"`
	Title        string        `gorm:"type:varchar(255);not null"`
	Description  string        `gorm:"type:text"`
	UnitPrice    int           `gorm:"type:integer"`
	Quantity     int           `gorm:"type:integer"`
	OrderDetails []OrderDetail `gorm:"foreignkey:ProductID"`
}
