package usecase

import (
	"github.com/Ma-Leal/to-do-list/internal/entity"
)

var statusMap = map[int]string{
	1: "to-do",
	2: "in-progress",
	3: "done",
}

type TaskUseCase struct {
	Repo entity.TaskRepositoryInterface
}

func NewTaskUseCase(repo entity.TaskRepositoryInterface) *TaskUseCase {
	return &TaskUseCase{Repo: repo}
}

func (t *TaskUseCase) CreateTask(desc string, statusID int) (entity.Task, error) {
	newTask := entity.NewTask(0, desc, statusMap[1])
	task, err := t.Repo.Save(*newTask)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (t *TaskUseCase) UpdateTaskDescription(id int, desc string) (entity.Task, error) {

	task, _ := t.Repo.GetByID(id)
	task.Update(desc)
	return task, t.Repo.Update(task)
}

func (t *TaskUseCase) UpdateTaskStatus(id int, statusID int) (entity.Task, error) {
	task, _ := t.Repo.GetByID(id)
	task.SetStatus(statusMap[statusID])
	return task, t.Repo.Update(task)
}

func (t *TaskUseCase) ListTasksByStatus(statusID int) ([]entity.Task, error) {
	return t.Repo.GetByStatus(statusMap[statusID])
}

func (t *TaskUseCase) ListTasks() ([]entity.Task, error) {
	return t.Repo.GetAll()
}

func (t *TaskUseCase) DeleteTask(id int) error {
	return t.Repo.Delete(id)
}
