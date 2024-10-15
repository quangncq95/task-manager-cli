package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type DeleteCommand struct {
	name string
}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{name: "delete"}
}

func (cmd *DeleteCommand) Exec() {
	deleteCommand := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	deleteCommand.Parse(os.Args[2:])

	if len(deleteCommand.Args()) < 1 {
		log.Fatalf("Task ID is missing !")
		os.Exit(1)
	}

	taskId := deleteCommand.Arg(0)

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
		os.Exit(1)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
		os.Exit(1)
	}

	var listTask []Task

	err = json.Unmarshal(data, &listTask)
	if err != nil {
		log.Fatalf("Error while parse json: %v", err)
		os.Exit(1)
	}

	for index, task := range listTask {
		if task.ID == taskId {
			list1 := listTask[:index]
			list2 := listTask[index+1:]

			newListTask := append(list1, list2...)

			err = saveListTask(file, newListTask)

			if err != nil {
				log.Fatalf("Error while saving file: %v", err)
				os.Exit(1)
			}

			break
		}
	}

	fmt.Printf("Delete task successfully ! Task ID : %v", taskId)

}

func (cmd *DeleteCommand) GetCommand() string {
	return cmd.name
}
