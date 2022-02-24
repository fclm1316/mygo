package Models

type whosyourdaddy struct {
	Id       int    `gorm:"type:Auto_increment;not null"`
	Ip       string `gorm:"size:15"`
	Username string `gorm:"size:10"`
	Password string `gorm:"size:10"`
}
