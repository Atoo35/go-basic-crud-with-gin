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
	routes.GET("/", middlewares.VerifyToken(), h.GetBooks)
}
