package controller

import (
	"gin-shop/service"
	"gin-shop/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ShopController struct {
}

func (sc *ShopController) Register(router *gin.Engine) {
	router.GET("/shops", getShop)
	router.GET("/search_shops", searchShop)
}

func searchShop(c *gin.Context) {
	//lo, la := strings.Split(c.Query("geohash"), ",")[0], strings.Split(c.Query("geohash"), ",")[1]
	keyword := c.Query("keyword")

	// 测试用默认值
	lo, la := 116.34, 40.34
	shops, err := service.SearchShop(lo, la, keyword)
	if err != nil {
		utils.Fail(c, "查询商铺列表失败")
		return
	}

	utils.Ok(c, shops)
}

func getShop(c *gin.Context) {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	lo, _ := strconv.ParseFloat(longitude, 32)
	la, _ := strconv.ParseFloat(latitude, 32)

	// 测试用默认值
	lo, la = 116.34, 40.34

	categories, err := service.GetShopList(lo, la)

	if err != nil {
		utils.Fail(c, "获取商铺列表失败")
		return
	}

	utils.Ok(c, categories)
}
