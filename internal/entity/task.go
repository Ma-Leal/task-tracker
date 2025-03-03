package entity

import (
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

const timeFormat = "02/01/06 15:04:05"

func NewTask(id int, description, status string) *Task {
	task := &Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now().Format(timeFormat),
	}
	task.SetStatus("to-do")
	return task
}

func (t *Task) Update(description string) {
	t.Description = description
	t.UpdatedAt = time.Now().Format(timeFormat)
}

func (t *Task) SetStatus(status string) {
	t.Status = status
	t.UpdatedAt = time.Now().Format(timeFormat)
}
