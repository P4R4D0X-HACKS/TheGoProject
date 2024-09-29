package main

import (
	"flag"
	"fmt"
	"log"
	"net"
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

func showInterfaces(){
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Error fetching interface", err)
	}

	for _, iface := range interfaces{
		if iface.Flags&net.FlagUp != 0 {
			fmt.Printf("%s: %s\n", iface.Name, iface.HardwareAddr)
		}
	}
}


func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -iface <interface>  Specify the interface for which you want to change the MAC address.")
	fmt.Println("  -newMac <address>   Specify the new MAC address to assign.")
	fmt.Println("Example:")
	fmt.Println("  go run main.go -iface eth0 -newMac 00:11:22:33:44:55")
	fmt.Println("Available network interfaces:")
	showInterfaces() // Show available interfaces in help
}

func main() {
	iface := flag.String("iface", "", "Interface for which you want to change the mac address")
	newMac := flag.String("newMac", "", "The custom mac address which you want to assig")
	help := flag.Bool("help", false, "Show help")

	if *help {
		displayHelp()
		return
	}

	if *iface == "" || *newMac == "" {
		fmt.Println("Error: Both -iface and -newMac flags must be provided.")
		displayHelp()
		return
	}

	flag.Parse()
	executeCommand("sudo", []string{"ifconfig",*iface,"down"})
	executeCommand("sudo", []string{"ifconfig",*iface,"hw", "ether", *newMac})
	executeCommand("sudo", []string{"ifconfig",*iface,"up"})
}