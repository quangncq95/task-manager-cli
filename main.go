package main

import (
	"log"
	"ncquang/task-manager/commands"
	"ncquang/task-manager/storage"
	"os"
	"path/filepath"
)

const fileName = "todo"
const dataFolder = "data"

func main() {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	absPath := filepath.Dir(execPath)
	storage, err := storage.NewFileStorage(filepath.Join(absPath, dataFolder), fileName)
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
