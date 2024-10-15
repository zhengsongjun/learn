package controllers

import "github.com/gin-gonic/gin"

type Order struct{}

func (o Order) GetOrder(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	ReturnSuccess(c, 200, "成功了", id+":"+name, 1)

}
