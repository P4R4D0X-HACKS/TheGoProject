# GoLang Structs and Pointers Example

This Go program demonstrates the use of structs and pointers. It explains how Go handles memory management using pointers and values, with a focus on struct manipulation.

## Overview

Go is a pass-by-value language, which means that when a variable is passed to a function, a copy of the value is made. To modify the original value within a function, you need to pass a pointer to the variable.


## Struct Initialization

Structs are user-defined types that group together different fields. Here's how to define and initialize structs.

```go
type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    contact   contactInfo
}
```

### Creating Structs

There are several ways to initialize a struct:

1. Using a struct literal:
```go
alex := person{"Alex", "Anderson"}
```

2. Using named fields:
```go
alex := person{firstName: "Alex", lastName: "Anderson"}
```

3. Using zero-value initialization (i.e., with empty fields):
```go
var alex person
```

4. You can also initialize nested structs:

```go
jim := person{
    firstName: "Jim",
    lastName:  "Party",
    contact: contactInfo{
        email:   "jim@gmail.com",
        zipCode: 741235,
    },
}
```

## Working with Pointers

Go allows you to pass pointers to a function to modify the original data. This is especially useful when working with value types like structs, where copying large amounts of data could be inefficient.

To work with pointers:

   -  Use the `&` operator to get the memory address of a variable.
   -  Use the `*` operator to dereference a pointer and access or modify the value at that memory address.

Example: Passing a Pointer to a Method

You can create a method that works with pointers to modify a struct's value:

```go
func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}
```

Here, `pointerToPerson *person` is just a type description meaning the method works with a pointer to a `person`. The `*` operator is used to dereference the pointer and access the original `firstName` field.

### Example in Action

In the main function, a new person struct is created and its name is updated using a pointer receiver method:

```go
func main() {
    jim := person{
        firstName: "Jim",
        lastName:  "Party",
        contact: contactInfo{
            email:   "jim@gmail.com",
            zipCode: 741235,
        },
    }

    jim.updateName("sam")
    jim.print()
}
```

This will update the `firstName` of the `jim` object to "sam" and print the updated values.

### Output

```go
{firstName:sam lastName:Party contact:{email:jim@gmail.com zipCode:741235}}
```

### Key Takeaways

-  **Value Types**: Integers, floats, strings, booleans, and structs are value types. If you want to modify them in functions, use pointers.
-  ** Reference Types**: Slices, maps, channels, functions, and pointers are reference types. You don't need to use pointers to modify them inside functions.
-   **Pointers**: Use * to dereference a pointer and & to get the memory address of a variable.

### Additional Notes

- The `*person` in the method signature signifies that we are using a pointer receiver. This allows you to modify the original struct in memory.
- Go automatically handles pointer dereferencing in most cases, so you can omit `*` when accessing struct fields via a pointer. However, explicitly using `*` can improve readability when working with complex code.

This example demonstrates fundamental concepts in Go for working with pointers and structs. It's essential for building efficient programs where memory manipulation is required.