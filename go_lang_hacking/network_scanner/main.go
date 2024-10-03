package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	IP := "scanme.nmap.org"

	for i := 1; i < 100; i++ {
		address := fmt.Sprintf(IP+":%d", i)
		
		// Adding a timeout for connection attempts to avoid hanging
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		
		if err == nil {
			fmt.Printf("[+] Connection Established... PORT %v\n", i)
			// Close the connection immediately after establishing it
			conn.Close()
		} else {
			fmt.Printf("[-] Connection Failed on PORT %v\n", i)
		}
	}
}
