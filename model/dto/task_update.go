package dto

type TaskUpdated struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
