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
	executeCommand("sudo", []string{"ifconfig","ens5","down"})
	executeCommand("sudo", []string{"ifconfig","ens5","hw", "ether", "00:11:22:33:44:55"})
	executeCommand("sudo", []string{"ifconfig","ens5","up"})
}