package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type MarkDoneCommand struct {
	name string
}

func NewMarkDoneCommand() *MarkDoneCommand {
	return &MarkDoneCommand{name: "mark-done"}
}

func (cmd *MarkDoneCommand) Exec() {
	markDoneCommand := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	markDoneCommand.Parse(os.Args[2:])

	if len(markDoneCommand.Args()) < 1 {
		log.Fatalf("Task ID is required !")
		os.Exit(1)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
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

	if len(markDoneCommand.Args()) < 1 {
		fmt.Printf("Task ID is required")
		os.Exit(1)
	}

	taskId := markDoneCommand.Arg(0)

	task := findTaskById(listTask, taskId)
	if task == nil {
		fmt.Printf("Don't found task with id : %v", taskId)
		return
	}

	updateTaskStatus(task, Done)

	err = saveListTask(file, listTask)
	if err != nil {
		log.Fatalf("Error %v", err)
		os.Exit(1)
	}

	fmt.Printf("Marked done task %v", taskId)

}

func updateTaskStatus(task *Task, status TaskStatus) {
	task.Status = status
	task.UpdateAt = time.Now().Format(time.DateTime)
}

func findTaskById(listTask []Task, id string) *Task {
	for index, task := range listTask {
		if task.ID == id {
			return &listTask[index]
		}
	}
	return nil
}

func (cmd *MarkDoneCommand) GetCommand() string {
	return cmd.name
}
