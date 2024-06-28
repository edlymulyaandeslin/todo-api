package dto

type TaskRequest struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
}
