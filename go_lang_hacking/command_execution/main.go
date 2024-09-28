package main

import (
	"log"
	"os"
	"os/exec"
)


func executeCommand(command string,args_arr []string)(err error) {

	args := args_arr
	cmd := exec.Command(command, args...) // Use a fetch a command with multiple arguments
	cmd.Stdout = os.Stdout // Link standard output
	cmd.Stderr = os.Stderr // Link standard error

	err = cmd.Run() // Run a command 
	if err != nil {
		log.Fatal("cmd.Run() failed ", err)
		return
	}
	return
}


func main() {
	command := "ls"
	executeCommand(command, []string{"-l"})
}