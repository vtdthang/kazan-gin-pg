package entities

import "database/sql"

// Order ...
type Order struct {
	ID           int            `gorm:"primary_key;auto_increment" json:"id"`
	UserID       int            `gorm:"index:idx_orders_user_id" json:"user_id"`
	User         User           `gorm:"foreignkey:UserID"`
	Phone        sql.NullString `gorm:"type:varchar(45)" json:"phone"`
	Address      sql.NullString `gorm:"type:text" json:"address"`
	Status       int8           `gorm:"type:smallint;not null" json:"status"`
	OrderDetails []OrderDetail  `gorm:"foreignkey:OrderID"`
}
