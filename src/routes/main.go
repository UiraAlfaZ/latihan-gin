package routes

import (
	controllers "latihan-gin/src/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		day := v1.Group("/day")
		{
			day.GET("/list", controllers.SelectAllDay)
			day.GET("/:id", controllers.SelectAllDay)
			day.POST("/create", controllers.InsertDay)
			day.PUT("/update/:id", controllers.UpdateDay)
			day.DELETE("/delete/:id", controllers.DeleteDay)
		}
		month := v1.Group("/month")
		{
			month.GET("/list", controllers.SelectAllMonth)
			month.GET("/:id", controllers.SelectAllDay)
			month.POST("/create", controllers.InsertMonth)
			month.PUT("/update/:id", controllers.UpdateMonth)
			month.DELETE("/delete/:id", controllers.DeleteMonth)
		}
		year := v1.Group("/year")
		{
			year.GET("/list", controllers.SelectAllYear)
			year.GET("/:id", controllers.SelectYear)
			year.POST("/create", controllers.InsertYear)
			year.PUT("/update/:id", controllers.UpdateYear)
			year.DELETE("/delete/:id", controllers.DeleteYear)
		}
	}
	router.Run(":8080")
}
