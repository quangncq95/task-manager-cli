package main

import (
	"log"
	"ncquang/task-manager/commands"
	"ncquang/task-manager/storage"
	"os"
)

const filePath = "todo.txt"

func main() {
	storage, err := storage.NewFileStorage(filePath)
	if err != nil {
		log.Fatalf("Init file storage failed : %v", err)
	}
	var listCommand []commands.CommandInterface
	add := commands.NewAddCommand(storage)
	listCommand = append(listCommand, add)
	delete := commands.NewDeleteCommand(storage)
	listCommand = append(listCommand, delete)
	markDone := commands.NewMarkDoneCommand(storage)
	listCommand = append(listCommand, markDone)
	list := commands.NewListCommand(storage)
	listCommand = append(listCommand, list)
	markInProgress := commands.NewMarkInProgressCommand(storage)
	listCommand = append(listCommand, markInProgress)
	updateCommand := commands.NewUpdateCommand(storage)
	listCommand = append(listCommand, updateCommand)

	var matchedCmd commands.CommandInterface = nil
	for _, cmd := range listCommand {
		if cmd.GetCommand() == os.Args[1] {
			matchedCmd = cmd
		}
	}

	if matchedCmd != nil {
		matchedCmd.Exec()
	} else {
		log.Fatalf("Command not found !")
	}
}
