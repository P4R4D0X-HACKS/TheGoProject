```Go
// package main

// import (
// 	"fmt"
// 	"net"
// 	"sync"
// )

// func main() {
// 	// Testing the connection

// 	fmt.Println("TCP Scanner")
// 	conn, err := net.Dial("tcp", "scanme.nmap.org:80")
// 	if err == nil {
// 		fmt.Println("Connection successful")
// 		fmt.Println(conn)
// 		fmt.Println("Local Address:", conn.LocalAddr().String())
// 		fmt.Println("Remote Adress:", conn.RemoteAddr().String())
// 		conn.Close()
// 	} else {
// 		fmt.Println("Connection unsuccessful")
// 	}

// 	//--------------------------------------------------------------------------------------------------------------------------

// 	// Testing the connection for multiple ports

// 	for i := 1; i <= 1024; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err == nil {
// 			fmt.Printf("Port %d is open\n", i)
// 			conn.Close()
// 		}
// 	}

// 	// --------------------------------------------------------------------------------------------------------------------------

// 	// Testing the connection for multiple ports with goroutines

// 	for i := 1; i <= 1024; i++ {
// 		go func(j int) {
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err == nil {
// 				fmt.Printf("Port %d is open\n", j)
// 				conn.Close()
// 			}
// 		}(i)
// 	}

// 	// --------------------------------------------------------------------------------------------------------------------------

// 	// Testing the connection for multiple ports with goroutines and wait group

// 	var wg sync.WaitGroup // Wait group is used to wait for the goroutines to finish, before the main function exits, this acts as a synchoronized counter
// 	for i := 1; i <= 1024; i++ {
// 		wg.Add(1) // Increment the wait group counter
// 		go func(j int) {
// 			defer wg.Done() // Decrement the wait group counter
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err == nil {
// 				fmt.Printf("Port %d is open\n", j)
// 				conn.Close()
// 			}
// 		}(i)
// 	}
// 	wg.Wait() // Wait for all the goroutines to finish

// 	//--------------------------------------------------------------------------------------------------------------------------

// }
```

