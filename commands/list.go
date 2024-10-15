package commands

import "fmt"

type ListCommand struct {
	name string
}

func NewListCommand() *ListCommand {
	return &ListCommand{name: "list"}
}

func (cmd *ListCommand) Exec() {
	fmt.Printf(cmd.name)
}

func (cmd *ListCommand) GetCommand() string{
	return cmd.name
}