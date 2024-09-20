
# Golang Concurrency: Using Goroutines and Channels

In this program, we explore key concurrency concepts in Golang, specifically Goroutines and Channels. The program checks the availability of a list of websites concurrently using these concepts, enabling it to efficiently handle multiple tasks simultaneously.

### Code Overview

We create a list of URLs and use goroutines to check each link concurrently. A channel is used to communicate between the goroutines and the main function, allowing the program to manage concurrency.

Here's the code:

```go 
package main

import ( 
    "fmt" 
    "net/http" 
    "time" 
)

func main() { 
    links := []string{ 
        "http://google.com", 
        "http://facebook.com", 
        "http://stackoverflow.com", 
        "http://golang.org", 
        "http://amazon.com", 
    }

    c := make(chan string)

    for _, link := range links { 
        go checkLink(link, c) 
    }

    for l := range c { 
        go func(link string) { 
            time.Sleep(5 * time.Second) 
            checkLink(link, c) 
        }(l) 
    } 
}

func checkLink(link string, c chan string) { 
    _, err := http.Get(link) 
    if err != nil { 
        fmt.Println(link, "might be down") 
        c <- link 
        return 
    }

    fmt.Println(link, "is up!") 
    c <- link 
} 
```

---

### Key Concepts

#### Goroutines 
Goroutines are lightweight threads managed by the Go runtime. They are started with the `go` keyword and allow functions to run concurrently.

In this program: 
```go 
go checkLink(link, c) 
``` 
This line starts a new goroutine that calls the `checkLink` function, which checks the status of the URL concurrently with other goroutines.

#### Channels 
Channels are a way to communicate between goroutines in Go. In this program, we create a channel of type `string` using:

```go 
c := make(chan string) 
```

Channels allow goroutines to send and receive values, enabling synchronization and communication. In this program, the `checkLink` function sends the link back to the channel once it's checked.

#### Anonymous Goroutine with Sleep 
The second loop in the main function runs another goroutine with a closure:

```go 
go func(link string) { 
    time.Sleep(5 * time.Second) 
    checkLink(link, c) 
}(l) 
```

This anonymous function delays execution by 5 seconds before calling `checkLink` again. This allows us to periodically check the status of the URLs.

#### Infinite Loop with Channels 
The infinite loop:

```go 
for l := range c { 
    // Goroutine logic 
} 
```

This loop listens for values from the channel `c`. Each time a link is sent from the `checkLink` function, the loop receives it and starts another goroutine to check the link again after 5 seconds.

---

### Program Execution Flow

1 . **Create Channel**: We first create a channel to communicate between goroutines. 
2 . **Start Goroutines**: For each link, a goroutine is started that checks the link's status via an HTTP GET request. 
3 . **Receive from Channel**: The program then listens for any messages coming through the channel (`c`), which it receives when a link is checked. 
4 . **Repeat After Delay**: For each received link, a new goroutine is spawned with a 5-second delay, continuing to check the status of the URL.

---

### Concurrency in Go

- **Goroutines**: Lightweight and managed by the Go runtime, making it possible to run thousands of them concurrently. 
- **Channels**: Essential for communicating between goroutines, avoiding race conditions, and synchronizing tasks.

This program demonstrates the power of Go's concurrency model by efficiently checking multiple website statuses without blocking the main thread.

---

This documentation explains the use of goroutines, channels, and how they interact within the program to manage concurrent tasks.