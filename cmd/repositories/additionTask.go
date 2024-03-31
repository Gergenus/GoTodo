package repositories

import (
	"log"
	"stepan/cmd/models"
	"stepan/cmd/storage"

	"github.com/sirupsen/logrus"
)

func TasksNew(task models.Tasks) {
	db := storage.GetDB()
	db.QueryRow("INSERT INTO tasks (title, description) VALUES ($1, $2)", task.Title, task.Description)
}

func TaskRight(id string) models.Tasks {
	info := models.Tasks{}
	db := storage.GetDB()
	db.QueryRow("SELECT title, description FROM tasks WHERE id = $1", id).Scan(&info.Title, &info.Description)
	return info
}

func Tasks() []models.Tasks {
	db := storage.GetDB()
	row, err := db.Query("SELECT title, description FROM tasks")
	if err != nil {
		log.Fatal("Ошибка при показе данных из базы")
	}
	varum := models.Tasks{}
	varum_list := []models.Tasks{}
	for row.Next() {
		err := row.Scan(&varum.Title, &varum.Description)
		if err != nil {
			log.Fatal("Ошибка показа таски")
		}
		varum_list = append(varum_list, varum)
	}
	return varum_list
}

func ReTasks(Uptask models.Tasks, id string) {
	db := storage.GetDB()
	_, err := db.Exec(`UPDATE tasks SET title = $1, description = $2 WHERE id = $3`, Uptask.Title, Uptask.Description, id)
	if err != nil {
		log.Fatal("Ошибка в переделке")
	}
}

func DeleteTask(id string) {
	db := storage.GetDB()
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("databse error: %s", err)
	}
}
