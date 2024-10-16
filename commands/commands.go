package commands

import (
	"encoding/json"
	"ncquang/task-manager/storage"
)

type CommandInterface interface {
	Exec()
	GetCommand() string
}

type Command struct {
	name    string
	storage storage.IStorage
}

type TaskStatus int

const (
	Todo TaskStatus = iota + 1
	InProgress
	Done
)

type Task struct {
	ID          string
	Description string
	Status      TaskStatus
	CreateAt    string
	UpdateAt    string
}

type TaskList []Task

func readListTask(storage storage.IStorage) (TaskList, error) {
	data, err := storage.Read()
	if err != nil {
		return nil, err
	}

	var listTodo []Task
	if len(data) > 0 {
		err = json.Unmarshal(data, &listTodo)
		if err != nil {
			return nil, err
		}
	}

	return listTodo, nil
}

func saveListTask(storage storage.IStorage, listTask []Task) error {
	listTodoBytes, err := json.Marshal(listTask)
	if err != nil {
		return err
	}

	storage.Write(listTodoBytes)
	if err != nil {
		return err
	}

	return nil
}

func getStatusString(status TaskStatus) string {
	if status == Todo {
		return "todo"
	} else if status == InProgress {
		return "in-progress"
	} else if status == Done {
		return "done"
	}

	return ""
}
