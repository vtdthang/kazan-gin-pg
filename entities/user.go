package entities

// User ...
type User struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string `gorm:"type:varchar(255);unique;not null" json:"email"`
	FirstName string `gorm:"type:varchar(255);unique;not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(255);unique;not null" json:"last_name"`
}
