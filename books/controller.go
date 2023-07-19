package books

import (
	"github.com/Atoo35/basic-crud/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/books")
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.POST("/add", middlewares.VerifyToken(), h.CreateBook)
	routes.PUT("/:id", middlewares.VerifyToken(), h.UpdateBook)
}
