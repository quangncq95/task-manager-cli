package main

import (
	"fmt"
	"ncquang/task-manager/commands"
	"os"
)

func main() {
	var listCommand []commands.CommandInterface
	list := commands.NewListCommand()
	listCommand = append(listCommand, list)
	add := commands.NewAddCommand()
	listCommand = append(listCommand, add)
	delete := commands.NewDeleteCommand()
	listCommand = append(listCommand, delete)
	markDone := commands.NewMarkDoneCommand()
	listCommand = append(listCommand, markDone)

	var matchedCmd commands.CommandInterface = nil
	for _, cmd := range listCommand {
		if cmd.GetCommand() == os.Args[1] {
			matchedCmd = cmd
		}
	}

	if matchedCmd != nil {
		matchedCmd.Exec()
	} else {
		fmt.Print("Command not found !")
	}
}
