package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Atoo35/basic-crud/auth"
	"github.com/Atoo35/basic-crud/books"
	"github.com/Atoo35/basic-crud/configurations"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

var (
	config configurations.Config
)

func init() {
	var err error
	config, err = configurations.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	configurations.Connect(&config)
	utils.SetSecretKey([]byte(config.JWT_SECRET))
}

func main() {
	router := gin.Default()
	books.RegisterRoutes(router, configurations.DB)
	auth.RegisterRoutes(router, configurations.DB)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	log.Fatal(router.Run(fmt.Sprintf(":%d", config.PORT)))
}
