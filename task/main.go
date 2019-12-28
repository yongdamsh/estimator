package task

import (
	"time"
)

type Feature struct {
	Id   int
	Name string
}

type Task struct {
	Id        int
	Feature   Feature
	Name      string
	Estimated string
	Elapsed   string
	CreatedAt time.Time
}

func NewHandler() *Handler {
	return &Handler{
		model: NewModel(),
	}
}
