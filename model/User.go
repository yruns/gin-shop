package model

type User struct {
	Id           int64   `gorm:"primary_key;column:id" json:"id"`
	Name         string  `gorm:"column:name" json:"name"`
	Password     string  `gorm:"column:password" json:"password"`
	Phone        string  `gorm:"column:phone" json:"phone"`
	RegisterTime int64   `gorm:"column:register_time" json:"register_time"`
	Avatar       string  `gorm:"column:avatar" json:"avatar"`
	Balance      float64 `gorm:"column:balance" json:"balance"`
	IsActive     int8    `gorm:"column:is_active" json:"is_active"`
	City         string  `gorm:"column:city" json:"city"`
}
