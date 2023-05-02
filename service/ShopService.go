package service

import (
	"gin-shop/model"
	"gin-shop/settings/db"
)

const DISTANCE_RANGE float64 = 100

func GetShopList(longitude, latitude float64) ([]model.Shop, error) {
	var shopList []model.Shop
	err := db.Mysql.Table("shop").
		Where("longitude < ? AND latitude < ? AND longitude > ? AND latitude > ?",
			longitude+DISTANCE_RANGE, latitude+DISTANCE_RANGE,
			longitude-DISTANCE_RANGE, latitude-DISTANCE_RANGE).
		Find(&shopList).Error
	return shopList, err
}

func SearchShop(longitude, latitude float64, keyword string) ([]model.Shop, error) {
	var shopList []model.Shop
	err := db.Mysql.Table("shop").
		Where("longitude < ? AND latitude < ? AND longitude > ? AND latitude > ?",
			longitude+DISTANCE_RANGE, latitude+DISTANCE_RANGE,
			longitude-DISTANCE_RANGE, latitude-DISTANCE_RANGE).
		Where("name LIKE ?", "%"+keyword+"%").
		Find(&shopList).Error
	return shopList, err
}
