package commands

import (
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"os"
	"time"
)

type MarkDoneCommand struct {
	cmd Command
}

func NewMarkDoneCommand(storage storage.IStorage) *MarkDoneCommand {
	return &MarkDoneCommand{cmd: Command{
		name:    "mark-done",
		storage: storage,
	}}
}

func (cmd *MarkDoneCommand) Exec() {
	markDoneCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	markDoneCommand.Parse(os.Args[2:])

	if len(markDoneCommand.Args()) < 1 {
		fmt.Printf("Task ID is required !")
		os.Exit(1)
	}

	listTask, err := readListTask(cmd.cmd.storage)
	if err != nil {
		log.Fatalf("Error ,%v", err)
	}

	taskId := markDoneCommand.Arg(0)

	task := listTask.findTaskById(taskId)
	if task == nil {
		fmt.Printf("Don't found task with id : %v", taskId)
		return
	}

	updateTaskStatus(task, Done)

	err = saveListTask(cmd.cmd.storage, listTask)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("Marked done task %v", taskId)

}

func updateTaskStatus(task *Task, status TaskStatus) {
	task.Status = status
	task.UpdateAt = time.Now().Format(time.DateTime)
}

func (listTask TaskList) findTaskById(id string) *Task {
	for index, task := range listTask {
		if task.ID == id {
			return &listTask[index]
		}
	}
	return nil
}

func (cmd *MarkDoneCommand) GetCommand() string {
	return cmd.cmd.name
}
