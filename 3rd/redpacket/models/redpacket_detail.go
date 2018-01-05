package models



type RedpacketDetail struct {
	Id  uint `gorm:"primary_key" json:"id"`
	RedpacketId uint `gorm:"index;not null" json:"redpacket_id"`
	UserId uint `gorm:"index;not null" json:"user_id"`
	Amount float64 `json:"amount"`
	//CreatedAt time.Time `json:"-"`
	//UpdatedAt time.Time `json:"-"`
}

