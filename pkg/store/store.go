package store

import (
	"hw1_backhard/internal/model"
)

type TaskStore interface {
	Create(task *model.Task) error
	Get(id string) (*model.Task, error)
	Update(task *model.Task) error
}
