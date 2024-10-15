package controllers

import (
	"go-gin-first/dao"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Sentence struct {
	Chinese     string    `json:"chinese"`
	Id          int64     `json:"id"`
	English     string    `json:"english"`
	LinkCourse  int64     `json:"linkCourse"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedAt   time.Time `json:"createAt"`
}

func (Sentence) TableName() string {
	return "sentences"
}

func (Sentence) GetAllSentences(c *gin.Context) {
	var sentences []Sentence
	query := dao.Db // 初始查询

	// 获取查询参数
	chinese := c.Query("chinese")
	english := c.Query("english")
	description := c.Query("desc")
	linkCourse := c.Query("linkCourse")

	if chinese != "" {
		query = query.Where("chinese LIKE ?", "%"+chinese+"%")
	}
	if english != "" {
		query = query.Where("english LIKE ?", "%"+english+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}
	if linkCourse != "" {
		query = query.Where("link_Course LIKE ?", "%"+linkCourse+"%")
	}

	if err := query.Find(&sentences).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "无法获取句子数据")
		return
	}

	count := int64(len(sentences))
	ReturnSuccess(c, http.StatusOK, "success", sentences, count)
}

func (Sentence) InsertSentence(c *gin.Context) {
	var newSentence Sentence
	if err := c.ShouldBindJSON(&newSentence); err != nil {
		ReturnError(c, http.StatusBadRequest, "请求数据格式错误")
		return
	}
	if err := dao.Db.Create(&newSentence).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "插入类别失败")
		return
	}
	ReturnSuccess(c, http.StatusCreated, "类别插入成功", newSentence, 0)
}

func (Sentence) DeleteSentence(c *gin.Context) {
	id := c.Param("id")
	var sentenceId int64
	if convertedId, err := strconv.ParseInt(id, 10, 64); err == nil {
		sentenceId = convertedId
	} else {
		ReturnError(c, http.StatusBadRequest, "无效的句子 ID")
		return
	}
	if err := dao.Db.Delete(&Sentence{}, sentenceId).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "删除句子失败")
		return
	}
	ReturnSuccess(c, http.StatusOK, "句子删除成功", nil, 0)
}

func (Sentence) UpdateSentence(c *gin.Context) {
	var updatedSentence Sentence
	id := c.Param("id")

	sentenceId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, "无效的句子 ID")
		return
	}

	if err := c.ShouldBindJSON(&updatedSentence); err != nil {
		ReturnError(c, http.StatusBadRequest, "请求数据格式错误")
		return
	}

	var existingSentence Sentence
	if err := dao.Db.First(&existingSentence, sentenceId).Error; err != nil {
		ReturnError(c, http.StatusNotFound, "句子未找到")
		return
	}

	if updatedSentence.Chinese != "" {
		existingSentence.Chinese = updatedSentence.Chinese
	}
	if updatedSentence.English != "" {
		existingSentence.English = updatedSentence.English
	}
	if updatedSentence.Description != "" {
		existingSentence.Description = updatedSentence.Description
	}

	if err := dao.Db.Save(&existingSentence).Error; err != nil {
		ReturnError(c, http.StatusInternalServerError, "更新句子失败")
		return
	}

	ReturnSuccess(c, http.StatusOK, "更新成功", existingSentence, 0)
}
