package service

import (
	"gin-shop/model"
	"gin-shop/settings/db"
)

func GetAllCategories() ([]model.FoodCategory, error) {
	var foodCategory []model.FoodCategory
	err := db.Mysql.Table("food_category").Find(&foodCategory).Error
	return foodCategory, err
}
