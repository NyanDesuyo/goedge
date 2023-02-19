package main

import (
	"github.com/gin-gonic/gin"

	"localhost/config"
	"localhost/controller"
)

func init() {
	config.LoadEnvronment()
}

func main() {
	r := gin.Default()

	r.GET("/", controller.Ping)

	todo := r.Group("/todo")
	{
		todo.POST("/", controller.TodoCreate)
		todo.GET("/", controller.TodoFindMany)
		todo.GET("/:id", controller.TodoFindOne)
		todo.PUT("/:id", controller.TodoUpdate)
		todo.DELETE("/:id", controller.TodoDelete)
	}

	account := r.Group("/account")
	{
		account.POST("/", controller.AccountCreate)
		account.GET("/", controller.AccountFindMany)
		account.GET("/:id", controller.AccountFindOne)
		account.PUT("/:id", controller.AccountUpdate)
		account.DELETE("/:id", controller.AccountDelete)
	}

	r.Run()
}
