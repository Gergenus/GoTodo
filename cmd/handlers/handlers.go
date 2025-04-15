package handlers

import (
	"net/http"
	"stepan/cmd/models"
	"stepan/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func TasksNew(c echo.Context) error {

	user := models.Tasks{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	repositories.TasksNew(user)
	return c.JSON(http.StatusCreated, user)
}

func TaskRight(c echo.Context) error {
	id := c.Param("id")
	data := repositories.TaskRight(id)
	return c.JSON(http.StatusOK, data)
}

func Tasks(c echo.Context) error {
	list := repositories.Tasks()
	return c.JSON(200, list)
}

func ReTasks(c echo.Context) error {
	id := c.Param("id")
	Update := models.Tasks{}
	c.Bind(&Update)
	repositories.ReTasks(Update, id)
	return c.JSON(200, Update)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	repositories.DeleteTask(id)
	return c.JSON(200, nil)
}
