package commands

import (
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"os"
)

type ListCommand struct {
	cmd Command
}

func NewListCommand(storage storage.IStorage) *ListCommand {
	return &ListCommand{cmd: Command{
		name:    "list",
		storage: storage,
	}}
}

func (cmd *ListCommand) Exec() {
	listCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	listCommand.Parse(os.Args[2:])

	listTask, err := readListTask(cmd.cmd.storage)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	for _, task := range listTask {
		status := ""
		if len(listCommand.Args()) > 0 {
			status = listCommand.Arg(0)
		}

		if len(status) > 0 {
			if status == "todo" && task.Status == Todo || status == "in-progress" && task.Status == InProgress || status == "done" && task.Status == Done {
				fmt.Printf("- %v - status:%v -ID:%v \n", task.Description, getStatusString(task.Status), task.ID)
			}
		} else {
			fmt.Printf("- %v - status:%v -ID:%v \n", task.Description, getStatusString(task.Status), task.ID)
		}

	}
}

func (cmd *ListCommand) GetCommand() string {
	return cmd.cmd.name
}
