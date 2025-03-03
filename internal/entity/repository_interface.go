package entity

type TaskRepositoryInterface interface {
	Save(Task)  (Task, error)
	GetAll() ([]Task, error)
	GetByID(id int) (Task, error)
	GetByStatus(status string) ([]Task, error)
	Update(task Task) error
	Delete(id int) error
}
