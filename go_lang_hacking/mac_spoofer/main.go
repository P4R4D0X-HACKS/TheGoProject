package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Validate MAC address format
func isValidMac(mac string) bool {
	// MAC address regex pattern
	macPattern := `^([0-9a-fA-F]{2}[:-]){5}([0-9a-fA-F]{2})$`
	re := regexp.MustCompile(macPattern)
	return re.MatchString(mac)
}

// Execute commands and check for errors
func executeCommand(command string, args_arr []string) error {
	args := args_arr
	cmd := exec.Command(command, args...) // Use a fetch a command with multiple arguments

	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout // Link standard output
	cmd.Stderr = &stderrBuf // Link standard error
	cmd.Stdin = os.Stdin // Link standard input

	err := cmd.Run() // Run a command 
	if err != nil {
		// Check if the error is due to 'ifconfig' not being found
		if exitError, ok := err.(*exec.ExitError); ok {
			if strings.Contains(stderrBuf.String(), "not found") {
				return fmt.Errorf("command not found")
			} else {
				// If not a 'not found' error, print the exit code
				fmt.Printf("Command exited with code: %d\n", exitError.ExitCode())
			}
		}
		return fmt.Errorf("cmd.Run() failed: %v", err) // Return a generic error
	}
	return nil
}

func showInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Error fetching interfaces: ", err)
	}

	for _, iface := range interfaces {
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
	iface := flag.String("iface", "", "Interface for which you want to change the MAC address")
	newMac := flag.String("newMac", "", "The custom MAC address you want to assign")
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

	// Validate the new MAC address format
	if !isValidMac(*newMac) {
		fmt.Println("Error: Invalid MAC address format.")
		displayHelp()
		return
	}

	showInterfaces() // Show the live interfaces 

	// Retrieve the original MAC address using the new function
	originalMac, err := getOriginalMac(*iface)
	if err != nil {
		log.Fatal(err)
	}

	// Execute commands and check for errors
	if err := executeCommand("sudo", []string{"ip", "link", "set", *iface, "down"}); err != nil {
		if strings.Contains(err.Error(), "command not found") {
			fmt.Println("Error: 'ip' command not found. Please install the iproute2 package.")
			return
		}
		log.Println(err) // Log the error without terminating the program
		return
	}

	if err := executeCommand("sudo", []string{"ip", "link", "set", *iface, "address", *newMac}); err != nil {
		if strings.Contains(err.Error(), "command not found") {
			fmt.Println("Error: 'ip' command not found. Please install the iproute2 package.")
			return
		}
		log.Println(err)
		return
	}

	if err := executeCommand("sudo", []string{"ip", "link", "set", *iface, "up"}); err != nil {
		if strings.Contains(err.Error(), "command not found") {
			fmt.Println("Error: 'ip' command not found. Please install the iproute2 package.")
			return
		}
		log.Println(err)
		return
	}

	fmt.Printf("\n%s\t\t%s\t\t%s\n", "Interface", "Original MAC", "New MAC")
	fmt.Printf("%s\t\t%s\t\t%s\n", *iface, originalMac, *newMac)
}
