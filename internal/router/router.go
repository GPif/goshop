package router

import (
	"goshop/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handler.Index)

	items := r.Group("/items")
	{
		items.GET("", handler.GetItems)
		items.POST("", handler.CreateItem)
		items.PUT("/:id", handler.UpdateItem)
		items.PATCH("/:id/toggle", handler.ToggleItem)
		items.DELETE("/:id", handler.DeleteItem)
	}

	return r
}
