package main

import (
	"stepan/cmd/handlers"
	"stepan/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()
	e.POST("/tasks", handlers.TasksNew)
	e.GET("tasks/:id", handlers.TaskRight)
	e.GET("/tasks", handlers.Tasks)
	e.PUT("tasks/:id", handlers.ReTasks)
	e.Logger.Fatal(e.Start(":8000"))
}
