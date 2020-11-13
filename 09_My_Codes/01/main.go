package main

import "github.com/gin-gonic/gin"

func main() {
	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	//Run attaches the router to a http.Server and starts listening and serving HTTP requests. It is a shortcut for http.ListenAndServe(addr, router)
	router.Run()
}
