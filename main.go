package main

import (
	"stepan/cmd/handlers"
	"stepan/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()
	api := e.Group("/tasks")
	{
		api.POST("/", handlers.TasksNew)
		api.GET("/:id", handlers.TaskRight)
		api.GET("/", handlers.Tasks)
		api.PUT("/:id", handlers.ReTasks)
		api.DELETE("/:id", handlers.DeleteTask)
	}
	e.Logger.Fatal(e.Start(":8000"))
}
