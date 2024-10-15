package commands

type CommandInterface interface {
	Exec()
	GetCommand() string
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

const filePath = "todo"