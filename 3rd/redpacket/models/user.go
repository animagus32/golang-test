package models


type User struct {
	Id  uint `gorm:"primary_key" json:"id"`
	Username string 		`gorm:"index;not null" json:"username"`
	Password string  	`gorm:"not null" json:"-"`
	Salt string 		`gorm:"not null" json:"-"`
	Balance float64 		`gorm:"not null;default:0" json:"balance"`

	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}

