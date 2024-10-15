package main

import (
	"go-gin-first/dao"
	"go-gin-first/router"
)

func main() {
	dao.Init()
	r := router.Router()
	r.Run(":9999")
}
