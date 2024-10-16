package commands

import (
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"os"
)

type MarkInProgressCommand struct {
	cmd Command
}

func NewMarkInProgressCommand(storage storage.IStorage) *MarkInProgressCommand {
	return &MarkInProgressCommand{cmd: Command{
		name:    "mark-in-progress",
		storage: storage,
	}}
}

func (cmd *MarkInProgressCommand) Exec() {
	markInProgressCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	markInProgressCommand.Parse(os.Args[2:])

	if len(markInProgressCommand.Args()) < 1 {
		fmt.Printf("Task ID is required !")
		os.Exit(1)
	}

	listTask, err := readListTask(cmd.cmd.storage)
	if err != nil {
		log.Fatalf("Error ,%v", err)
	}

	taskId := markInProgressCommand.Arg(0)

	task := listTask.findTaskById(taskId)
	if task == nil {
		fmt.Printf("Don't found task with id : %v", taskId)
		return
	}

	updateTaskStatus(task, InProgress)

	err = saveListTask(cmd.cmd.storage, listTask)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("Marked in-progress task %v", taskId)

}

func (cmd *MarkInProgressCommand) GetCommand() string {
	return cmd.cmd.name
}
