package models

import "time"

type Redpacket struct{
	Id  uint `gorm:"primary_key" json:"id"`
	UserId int `gorm:"index;not null" json:"user_id"`
	Amount float64 `gorm:"not null" json:"amount"`
	Secret string `gorm:"index;not null" json:"secret"`
	Status int8 `gorm:"default:0" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
