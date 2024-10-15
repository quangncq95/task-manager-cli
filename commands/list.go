package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type ListCommand struct {
	name string
}

func NewListCommand() *ListCommand {
	return &ListCommand{name: "list"}
}

func (cmd *ListCommand) Exec() {
	listCommand := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	listCommand.Parse(os.Args[2:])

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Error while opening file : %v", err)
		os.Exit(1)
	}

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
		os.Exit(1)
	}

	var listTask []Task
	err = json.Unmarshal(data, &listTask)
	if err != nil {
		log.Fatalf("Error while parsing json: %v", err)
		os.Exit(1)
	}

	for _, task := range listTask {
		status := ""
		if len(listCommand.Args()) > 0 {
			status = listCommand.Arg(0)
		}

		if len(status) > 0 {
			if status == "todo" && task.Status == Todo || status == "in-progress" && task.Status == InProgress || status == "done" && task.Status == Done {
				fmt.Printf("- %v - status:%v \n", task.Description, getStatusString(task.Status))
			}
		} else {
			fmt.Printf("- %v - status:%v \n", task.Description, getStatusString(task.Status))
		}

	}
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

func (cmd *ListCommand) GetCommand() string {
	return cmd.name
}
