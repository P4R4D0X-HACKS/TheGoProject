package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "scanme.nmap.org"

	for i := 1; i < 100; i++ {
		address := fmt.Sprintf(IP+":%d", i)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("[+] Connection Established... PORT %v\n", i)
			// Close the connection after use
			conn.Close()
		} else {
			fmt.Printf("[-] Connection Failed on PORT %v\n", i)
		}
	}
}
