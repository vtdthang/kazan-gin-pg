package models

type OrderCreateRequest struct {
}

type OrderListSql struct {
	ID      int
	UserID  int
	Phone   string
	Address string
	Status  int8
}
