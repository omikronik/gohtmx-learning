package todo

import (
	"time"

	"github.com/google/uuid"
)

type TodoItemRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoItem struct {
	Id        string
	Title     string
	Content   string
	Complete  bool
	CreatedOn time.Time
}

func NewToDoItem(title string, content string, complete bool) TodoItem {
	return TodoItem{
		Id:        uuid.NewString(),
		Title:     title,
		Content:   content,
		Complete:  complete,
		CreatedOn: time.Now(),
	}
}
