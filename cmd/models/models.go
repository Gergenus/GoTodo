package models

type Tasks struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type TasksWithNum struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
