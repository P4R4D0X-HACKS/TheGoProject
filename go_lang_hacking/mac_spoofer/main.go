package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)


func executeCommand(command string,args_arr []string)(err error) {

	args := args_arr
	cmd := exec.Command(command, args...) // Use a fetch a command with multiple arguments
	cmd.Stdout = os.Stdout // Link standard output
	cmd.Stderr = os.Stderr // Link standard error

	err = cmd.Run() // Run a command 
	if err != nil {
		// Check if the error is due to 'ifconfig' not being found
		if exitError, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(exitError.Stderr), "not found") {
				fmt.Println("Error: 'ifconfig' command not found. Please install the net-tools package.")
				fmt.Println("You can do this by running: sudo apt-get install net-tools")
				return nil // Return nil to prevent logging a fatal error
			}
		}
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

func getOriginalMac(interfaceName string) (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, ifaceDetail := range interfaces {
		if ifaceDetail.Name == interfaceName && ifaceDetail.Flags&net.FlagUp != 0 { // Check if the interface is active
			return ifaceDetail.HardwareAddr.String(), nil // Return the MAC address
		}
	}

	return "", fmt.Errorf("interface %s not found or not active", interfaceName)
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
	flag.Parse()

	if *help {
		displayHelp()
		return
	}

	if *iface == "" || *newMac == "" {
		fmt.Println("Error: Both -iface and -newMac flags must be provided.")
		displayHelp()
		return
	}

	showInterfaces() // show the live interfaces 
	
	// Retrieve the original MAC address using the new function
	originalMac, err := getOriginalMac(*iface)
	if err != nil {
		log.Fatal(err)
	}

	executeCommand("sudo", []string{"ifconfig",*iface,"down"})
	executeCommand("sudo", []string{"ifconfig",*iface,"hw", "ether", *newMac})
	executeCommand("sudo", []string{"ifconfig",*iface,"up"})

	fmt.Printf("\n%s\t\t%s\t\t%s\n", "Interface", "Original MAC", "New MAC")
	fmt.Printf("%s\t\t%s\t\t%s\n", *iface, originalMac, *newMac)
}