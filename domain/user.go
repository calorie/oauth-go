package domain

type User struct {
	Id       string `gorm:"primaryKey"`
	Email    string `gorm:"not null;type:varchar(320)"`
	Password string `gorm:"not null"`
}
