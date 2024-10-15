package commands

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"ncquang/task-manager/utils"
	"os"
	"time"
)

type AddCommand struct {
	name string
}

func NewAddCommand() *AddCommand {
	return &AddCommand{name: "add"}
}

func (cmd *AddCommand) Exec() {
	addCommand := flag.NewFlagSet(cmd.name, flag.ExitOnError)
	addCommand.Parse(os.Args[2:])
	if len(addCommand.Args()) < 1 {
		fmt.Println("Task description is missing !")
		os.Exit(1)
	} else {
		var file *os.File
		file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

		if os.IsNotExist(err) {
			file, err = os.Create(filePath)

			if err != nil {
				log.Fatalf("Error %v", err)
			}
		}

		defer file.Close()

		newTask := createNewTask(addCommand.Arg(0))
		err = saveNewTask(file, newTask)

		if err != nil {
			log.Fatalf("Error : %v", err)
		}

		fmt.Printf("Task added successfully ! (ID:%s)", newTask.ID)

		return
	}
}

func saveNewTask(file *os.File, newTask *Task) error {
	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	var listTodo []Task
	if len(data) > 0 {
		err = json.Unmarshal(data, &listTodo)
		if err != nil {
			return err
		}
	}

	listTodo = append(listTodo, *newTask)

	err = saveListTask(file, listTodo)
	if err != nil {
		return err
	}
	return nil
}

func saveListTask(file *os.File, listTask []Task) error {
	listTodoBytes, err := json.Marshal(listTask)
	if err != nil {
		return err
	}

	file.Seek(0, 0)
	file.Truncate(int64(len(listTodoBytes)))

	_, err = file.Write(listTodoBytes)
	if err != nil {
		return err
	}

	return nil
}

func createNewTask(taskDescription string) *Task {
	id := utils.GenerateTimestampBasedID()
	newTask := Task{
		ID:          id,
		Description: taskDescription,
		Status:      Todo,
		CreateAt:    time.Now().String(),
		UpdateAt:    time.Now().String(),
	}

	return &newTask
}

func (cmd *AddCommand) GetCommand() string {
	return cmd.name
}
