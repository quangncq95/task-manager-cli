package commands

import (
	"flag"
	"fmt"
	"log"
	"ncquang/task-manager/storage"
	"ncquang/task-manager/utils"
	"os"
	"time"
)

type AddCommand struct {
	cmd Command
}

func NewAddCommand(storage storage.IStorage) *AddCommand {
	return &AddCommand{cmd: Command{
		name:    "add",
		storage: storage,
	}}
}

func (cmd *AddCommand) Exec() {
	addCommand := flag.NewFlagSet(cmd.cmd.name, flag.ExitOnError)
	addCommand.Parse(os.Args[2:])
	if len(addCommand.Args()) < 1 {
		fmt.Println("Task description is missing !")
		os.Exit(1)
	} else {
		newTask := createNewTask(addCommand.Arg(0))

		err := saveNewTask(cmd.cmd.storage, newTask)

		if err != nil {
			log.Fatalf("Error : %v", err)
		}

		fmt.Printf("Task added successfully ! (ID:%s)", newTask.ID)

		return
	}
}

func saveNewTask(storage storage.IStorage, newTask *Task) error {
	listTodo, err := readListTask(storage)
	if err != nil {
		return err
	}

	listTodo = append(listTodo, *newTask)

	err = saveListTask(storage, listTodo)
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
		CreateAt:    time.Now().Format(time.DateTime),
		UpdateAt:    time.Now().Format(time.DateTime),
	}

	return &newTask
}

func (cmd *AddCommand) GetCommand() string {
	return cmd.cmd.name
}
