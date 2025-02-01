package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	// Print a welcome message to the console
	fmt.Println("Welcome to Go programming!")

	assign()
	sliceMap()
	pointers()

	guy := &Person{}
	guy.Name = "John"
	guy.SayHello()
}


// Example 1
// assign is a function that initializes two variables, x and z, and prints their values.
// x is a string with the value "Hello, World".
// z is an integer with the value 42.
func assign(){
	var x = "Hello, World"
	z := int(42)
	fmt.Println(x)
	fmt.Println(z)
}


// Example 2
// sliceMap is a function that demonstrates the creation and usage of a slice and a map.
// It initializes an empty slice of strings and appends the string "Hello, World" to it,
// then prints the slice. It also initializes an empty map with string keys and integer values,
// assigns the value 42 to the key "key", and prints the map.
func sliceMap(){
	var s = make([]string, 0)
	s = append(s, "Hello, World")
	fmt.Println(s)

	var m = make(map[string]int)
	m["key"] = 42 
	m["key2"] = 43
	fmt.Println(m)
}

// Example 3

func pointers(){
	var a int = 10
    var b *int = &a

    fmt.Println("Address of a:", &a)
    fmt.Println("Value of a:", a)
    fmt.Println("Address stored in pointer b:", b)
    fmt.Println("Value pointed to by b:", *b)

	var count = int(42)
	ptr := &count
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]\tValue Pointed By Ptr[", *ptr, "]")
	*ptr = 100
	fmt.Println(count)

	c := 12
	d := c

	fmt.Println(d)
}


// Example 4
// struct is a function that demonstrates the creation and usage of a struct.

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}


