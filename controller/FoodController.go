package controller

import (
	"gin-shop/service"
	"gin-shop/utils"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fc *FoodCategoryController) Register(router *gin.Engine) {
	router.GET("/food_category", getAllCategories)
}

func getAllCategories(c *gin.Context) {
	categories, err := service.GetAllCategories()
	if err != nil {
		utils.Fail(c, "获取食品分类失败")
		return
	}
	utils.Ok(c, categories)
}
