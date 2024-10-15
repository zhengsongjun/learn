package controllers

import (
	"go-gin-first/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Category struct {
	Name      string    `json:"name"`
	Id        int64     `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createAt"`
}

func (Category) TableName() string {
	return "category"
}

func (Category) GetAllCategory(c *gin.Context) {
	var categories []Category
	if err := dao.Db.Find(&categories).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "无法获取类别数据")
		return
	}
	count := int64(len(categories))
	ReturnSuccess(c, http.StatusOK, "success", categories, count)
}

func (Category) InsertCategory(c *gin.Context) {
	var newCategory Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		ReturnError(c, http.StatusBadRequest, "请求数据格式错误")
		return
	}
	if err := dao.Db.Create(&newCategory).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "插入类别失败")
		return
	}
	ReturnSuccess(c, http.StatusCreated, "类别插入成功", newCategory, 0)
}
