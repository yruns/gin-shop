package main

import (
	"gin-shop/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/cloud_shop"+
		"?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 生成表时指定表名
	//db.Table("food_category").AutoMigrate(&model.FoodCategory{})
	db.Table("shop").AutoMigrate(&model.Shop{})

}
