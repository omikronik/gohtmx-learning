package todo

import (
	"time"
)

type TodoItem struct {
	id        string
	title     string
	content   string
	complete  bool
	createdOn time.Time
}
