package controllers

import (
	"go-gin-first/dao"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Name         string    `json:"name"`
	Id           int64     `json:"id"`
	LinkCategory int64     `json:"linkCategory"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CreatedAt    time.Time `json:"createAt"`
}

func (Course) TableName() string {
	return "course"
}

func (Course) InsertCourse(c *gin.Context) {
	var newCourse Course

	// 绑定 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// 插入新记录
	if err := dao.Db.Create(&newCourse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "插入课程失败"})
		return
	}

	// 成功插入后返回课程信息
	c.JSON(http.StatusCreated, gin.H{"message": "课程插入成功", "data": newCourse})
}

// 查询所有课程
func (Course) GetAllCourses(c *gin.Context) {
	var courses []Course

	// 执行查询
	if err := dao.Db.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取课程数据"})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": courses})
}

func (Course) GetCourseByCategoryId(c *gin.Context) {
	categoryIdStr := c.Param("id")

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的分类 ID"})
		return
	}

	var courses []Course

	if err := dao.Db.Where("link_categote = ?", categoryId).Find(&courses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到相关课程"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": courses})
}

func (Course) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	courseId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的课程 ID"})
		return
	}

	var existingCourse Course

	if err := dao.Db.First(&existingCourse, courseId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "课程未找到"})
		return
	}

	var updatedCourse Course
	if err := c.ShouldBindJSON(&updatedCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		return
	}

	if updatedCourse.Name != "" {
		existingCourse.Name = updatedCourse.Name
	}
	if updatedCourse.LinkCategory != 0 {
		existingCourse.LinkCategory = updatedCourse.LinkCategory
	}

	if err := dao.Db.Save(&existingCourse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新课程失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "课程更新成功", "data": existingCourse})
}

func (Course) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	courseId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的课程 ID"})
		return
	}

	if err := dao.Db.Delete(&Course{}, courseId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除课程失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "课程删除成功"})
}
