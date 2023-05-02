package model

type FoodCategory struct {
	Id          int64  `gorm:"column:id;primary_key" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	ImageUrl    string `gorm:"column:image_url" json:"image_url"`
	LinkUrl     string `gorm:"column:link_url" json:"link_url"`
	IsInServing bool   `gorm:"column:is_in_serving" json:"is_in_serving"`
}
