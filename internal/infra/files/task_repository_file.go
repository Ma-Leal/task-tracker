package files

import (
	"encoding/json"
	"os"

	"github.com/Ma-Leal/to-do-list/internal/entity"
)

type TaskRepositoryFile struct {
	path   string
	lastID int
}

func NewTaskRepositoryFile(path string) *TaskRepositoryFile {
	repo := &TaskRepositoryFile{
		path: path,
	}
	repo.initialize()
	return repo
}

func (t *TaskRepositoryFile) initialize() {
	tasks, _ := t.GetAll()
	var maxID int

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	t.lastID = maxID

}

func (t *TaskRepositoryFile) Save(task entity.Task) (entity.Task, error) {

	t.lastID++
	task.ID = t.lastID

	tasks, _ := t.GetAll()
	tasks = append(tasks, task)

	return task, t.WriteFile(tasks)
}

func (t *TaskRepositoryFile) GetAll() ([]entity.Task, error) {

	dat, err := os.ReadFile(t.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []entity.Task{}, nil
		}
		return nil, err
	}

	var tasks []entity.Task
	err = json.Unmarshal(dat, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepositoryFile) Update(updatedTask entity.Task) error {

	tasks, _ := t.GetAll()
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			return t.WriteFile(tasks)
		}
	}
	return os.ErrNotExist
}

func (t *TaskRepositoryFile) Delete(id int) error {
	tasks, _ := t.GetAll()
	var newTasks []entity.Task

	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	return t.WriteFile(newTasks)
}

func (t *TaskRepositoryFile) GetByID(id int) (entity.Task, error) {
	tasks, err := t.GetAll()
	if err != nil {
		return entity.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return entity.Task{}, nil
}

func (t *TaskRepositoryFile) GetByStatus(status string) ([]entity.Task, error) {

	tasks, _ := t.GetAll()
	var newTasks []entity.Task

	for _, task := range tasks {
		if task.Status == status {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks, nil
}

func (t *TaskRepositoryFile) WriteFile(tasks []entity.Task) error {
	tasks_json, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(t.path, tasks_json, 0644)
}
