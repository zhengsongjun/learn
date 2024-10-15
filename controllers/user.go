package controllers

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func (u User) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, "success", id+":"+name, 1)
}

func (u User) GetUserList(c *gin.Context) {
	ReturnError(c, 404, "没有任何信息")
}
func (u User) PostUser(c *gin.Context) {
	params := make(map[string]interface{})
	err := c.Bind(&params)
	if err == nil {
		// 绑定成功时返回成功
		ReturnSuccess(c, 200, params["name"], params["id"], 10)
	} else {
		// 绑定失败时返回错误
		ReturnError(c, 500, gin.H{"err": err.Error()})
	}
}
