package commands

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"os"
)

type DeleteCommand struct {
	cmd Command
}

func NewDeleteCommand(storage storage.IStorage) *DeleteCommand {
	return &DeleteCommand{cmd: Command{
		name:    "delete",
		storage: storage,
	}}
}

func (cmd *DeleteCommand) Exec() {
	deleteCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	deleteCommand.Parse(os.Args[2:])

	if len(deleteCommand.Args()) < 1 {
		fmt.Printf("Task ID is missing !")
		os.Exit(1)
	}

	taskId := deleteCommand.Arg(0)

	listTask, err := readListTask(cmd.cmd.storage)

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	err = listTask.removeTaskById(taskId)
	if err != nil {
		log.Fatalf("Error :%v", err)
	}

	err = saveListTask(cmd.cmd.storage, listTask)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	fmt.Printf("Delete task successfully ! Task ID : %v", taskId)

}

func (listTask *TaskList) removeTaskById(taskID string) error {
	for index, task := range *listTask {
		if task.ID == taskID {
			list1 := (*listTask)[:index]
			list2 := (*listTask)[index+1:]

			*listTask = append(list1, list2...)

			return nil
		}
	}
	return errors.New("Task not found")
}

func (cmd *DeleteCommand) GetCommand() string {
	return cmd.cmd.name
}
