package order

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vtdthang/kazan-gin-pg/models"
)

// OrderRepository ...
type OrderRepository struct {
	DB *gorm.DB
}

// ProvideOrderRepository ...
func ProvideOrderRepository(DB *gorm.DB) OrderRepository {
	return OrderRepository{DB: DB}
}

// GetOrdersByUserID ...
func (o *OrderRepository) GetOrdersByUserID(userID int) {
	// var orders []entities.Order
	// o.DB.
	// 	Preload("User").
	// 	Preload("OrderDetails").
	// 	Preload("OrderDetails.Product").
	// 	Where("user_id = ?", userID).
	// 	Limit(1).
	// 	Offset(1).
	// 	Find(&orders)

	// for _, order := range orders {
	// 	fmt.Printf("%+v\n", order.User)
	// }

	orders2 := make([]models.OrderListSql, 0)
	sqlQuery := `
		SELECT
			o.id, o.user_id, o.Phone, o.Address, o.Status
		FROM
			orders o
		INNER JOIN
			order_details od ON o.id = od.order_id
		INNER JOIN
			products p ON od.product_id = p.id
		INNER JOIN
			users u ON o.user_id = u.id
		WHERE
			o.user_id = ?
	`
	rows, _ := o.DB.Raw(sqlQuery, userID).Rows()

	defer rows.Close()

	for rows.Next() {
		o := models.OrderListSql{}
		err := rows.Scan(&o.ID, &o.UserID, &o.Phone, &o.Address, &o.Status)
		if err != nil {
			fmt.Println(err)
		}

		orders2 = append(orders2, o)
	}

	for _, bk := range orders2 {
		fmt.Printf("%+v\n", bk)
	}
}
