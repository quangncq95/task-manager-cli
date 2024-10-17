package commands

import (
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"os"
)

type UpdateCommand struct {
	cmd Command
}

func NewUpdateCommand(storage storage.IStorage) *UpdateCommand {
	return &UpdateCommand{cmd: Command{
		name:    "update",
		storage: storage,
	}}
}

func (cmd *UpdateCommand) Exec() {
	updateCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	updateCommand.Parse(os.Args[2:])

	if len(updateCommand.Args()) < 2 {
		fmt.Printf("Please pass task id and new task's description")
		os.Exit(1)
	}

	listTask, err := readListTask(cmd.cmd.storage)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	task := listTask.findTaskById(updateCommand.Arg(0))
	if task == nil {
		fmt.Printf("Task not found !")
		os.Exit(1)
	}

	task.Description = updateCommand.Arg(1)

	saveListTask(cmd.cmd.storage, listTask)

	fmt.Printf("Update task successfully !")

}

func (cmd *UpdateCommand) GetCommand() string {
	return cmd.cmd.name
}
