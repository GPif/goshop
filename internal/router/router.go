package router

import (
	"goshop/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

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
