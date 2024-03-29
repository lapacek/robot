package main

import (
	"fmt"
	"os"
)

func main() {
	command := NewRobotCommand()
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
