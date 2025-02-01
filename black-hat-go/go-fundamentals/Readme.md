## Data Types

1. Primitive Types 
    - strings 
    - numbers 
    - booleans

2. Complex Data types > also know as user defiend structurs 
    - combination of one or more primitive data types 



> Primitive data types include -> bool, string, int, int[8,16,32,64], uint[8,16,32,64], uintptr, byte, rune, float[32,64], complex[64,128]

**Assigning a variable //check for example 1 in the main.go file**

### Slices and Maps (Complex data type)

- *Slices* are like arrays that we can dynamically resize and pass to functions
- *Maps* are associative arryas, unordered list of keys/value pair 

**Assigning a Slice and Map //check example 2**

### Pointers, Structs, and Interfaces

- A pointer points to a particular area in memory and allows us to retrieve the value stored there, the & operator can be used to retrieve the memory of some variable

**Pointer Example**

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b *int = &a

    fmt.Println("Address of a:", &a)
    fmt.Println("Value of a:", a)
    fmt.Println("Address stored in pointer b:", b)
    fmt.Println("Value pointed to by b:", *b)
}
```

In this example, `a` is an integer variable, and `b` is a pointer to an integer. The `&` operator is used to get the memory address of `a`, which is then stored in `b`. The `*` operator is used to access the value stored at the address pointed to by `b`.


- We can use the struct type to define new data types by specifiying the type's associated feilds and methods

**Struct Example**

```go
package main

import "fmt"

// Define a struct type
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the struct
    person1 := Person{Name: "Alice", Age: 30}

    // Access struct fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
}
```

In this example, we define a `Person` struct with two fields: `Name` (a string) and `Age` (an integer). We then create an instance of the `Person` struct and access its fields.


### Control Structures

Go have slightly fewer control structure 

```Go

if x == 1 {
    fmt.Println("X is equal to 1")
} else {
    fmt.println("X is not equal to 1")
}
```

```Go

switch x {
    case "foo":
        fmt.Println("Found foo")
    case "bar":
        fmt.Println("Found bar")
    default:
        fmt.Println("Default case")
}
```

```Go
func foo(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Println("I'm an integer")
        case string:
            fmt.Println("I'm a string")
        default:
            fmt.Println("Unknown type!")
    }
}
```

### Loop

"""
This script demonstrates the use of loops in Python with examples.

Loops are used to execute a block of code repeatedly. Python provides two types of loops:
1. `for` loop: Iterates over a sequence (such as a list, tuple, dictionary, set, or string).
2. `while` loop: Repeats as long as a condition is true.

Examples:

1. `for` loop:
    ```python
    fruits = ["apple", "banana", "cherry"]
    for fruit in fruits:
        print(fruit)
    ```
    This loop iterates over each item in the `fruits` list and prints it.

2. `while` loop:
    ```python
    count = 1
    while count <= 5:
        print(count)
        count += 1
    ```
    This loop prints numbers from 1 to 5. It continues to execute as long as `count` is less than or equal to 5.
"""

```Go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

```Go
nums := []int{2,4,5,6}
for idx, val := range nums {
    fmt.Println(idx, val)
}
```

### Concurrency in Go

Concurrency is the ability of a program to make progress on multiple tasks simultaneously. Go has rich support for concurrency using goroutines and channels.

#### Goroutines

A goroutine is a lightweight thread managed by the Go runtime. Goroutines are functions or methods that run concurrently with other functions or methods. They are created using the `go` keyword.

**Goroutine Example**

```go
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go printNumbers() // Start printNumbers in a new goroutine
    fmt.Println("Goroutine started")
    time.Sleep(6 * time.Second) // Wait for the goroutine to finish
    fmt.Println("Main function finished")
}
```

In this example, the `printNumbers` function runs concurrently with the main function. The `time.Sleep` calls are used to simulate work and ensure the main function waits for the goroutine to complete.

#### Channels

Channels are used to communicate between goroutines. They provide a way for one goroutine to send data to another goroutine.

**Channel Example**

```go
package main

import "fmt"

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total // Send total to channel c
}

func main() {
    a := []int{1, 2, 3, 4, 5}
    c := make(chan int)
    go sum(a, c) // Start sum in a new goroutine
    result := <-c // Receive result from channel c
    fmt.Println("Sum:", result)
}
```

In this example, the `sum` function calculates the sum of a slice of integers and sends the result to a channel. The main function receives the result from the channel and prints it.

Concurrency in Go allows you to write efficient and scalable programs by leveraging goroutines and channels to perform multiple tasks simultaneously.

### Error Handling

Error handling in Go is done using the built-in `error` type. Functions that can encounter errors typically return an `error` as the last return value. If the error is `nil`, it means no error occurred; otherwise, it contains information about the error.

**Error Handling Example**

```go
package main

import (
    "errors"
    "fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

In this example, the `divide` function returns an error if the divisor is zero. The `main` function checks for the error and handles it appropriately.

Go also provides the `fmt.Errorf` function to format error messages.

**Using fmt.Errorf**

```go
package main

import (
    "fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide %f by zero", a)
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

In this example, `fmt.Errorf` is used to create a formatted error message.

Proper error handling is crucial for writing robust and reliable Go programs. Always check for errors and handle them appropriately.