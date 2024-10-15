package router

import (
	"go-gin-first/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	category := r.Group("/category")
	{
		category.GET("/list", controllers.Category{}.GetAllCategory)
		category.POST("/list", controllers.Category{}.InsertCategory)
	}

	course := r.Group("/course")
	{
		course.GET("", controllers.Course{}.GetAllCourses)
		course.POST("", controllers.Course{}.InsertCourse)
		course.DELETE("", controllers.Course{}.DeleteCourse)
		course.PUT("", controllers.Course{}.UpdateCourse)
	}

	sentence := r.Group("/sentence")
	{
		sentence.GET("/all", controllers.Sentence{}.GetAllSentences)
		sentence.POST("/insert", controllers.Sentence{}.InsertSentence)
		sentence.DELETE("/:id", controllers.Sentence{}.DeleteSentence)
		sentence.PUT("/:id", controllers.Sentence{}.UpdateSentence)
	}
	// user := r.Group("/user")
	// {
	// 	user.GET("/info/:id/:name", controllers.User{}.GetUserInfo)
	// 	user.GET("/list", controllers.User{}.GetUserList)
	// 	user.POST("/list", controllers.User{}.PostUser)

	// 	user.DELETE("/list", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "删除user")
	// 	})

	// 	user.PUT("/list", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "修改user")
	// 	})
	// }

	// order := r.Group("/order")
	// {
	// 	order.GET("", controllers.Order{}.GetOrder)
	// }
	return r
}
