package entities

// Category ...
type Category struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
}
