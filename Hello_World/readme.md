## What did I learn from this program !!

#### Some questions that came in my mind 

1. How do we run the code in our project ?
2. What does 'package main' mean ?
3. What does 'import "fmt"' mean ?
4. What's the 'func' thing ? 
5. How is the main.go file oragnized ?

#### Answers: 

1. To run the code in our project we write `go run` 

**Here are some go command to run the project**

- `go build` - compiles a bunch of go source code file 
- `go run` - compiles and executes one or two files 
- `go fmt` - formats all the code in each file in the crrent directory 
- `go install` - Compiles and "installs" a package. 
- `go get` - downloads the raw source code of someone else's package 
- `go test` - runs any test assouciated with the current project 

---

2. Package == Project == Workspace. Like if we are working on a single app right now we are basically working on one package. Package can have many related file inside of it. **The only requirement that we have is we have to mention the package name one every file start.**

    There are two types of packages : 
    
    a. Executable type - Generates a file that can run 

    b. Reusable type - Good place to put reusable logic 

So if we are writing `package main` this means when we write go build it will make a executable file but if we write anyother thing except main, it will not generate any executable. *If we want to have a executable package it should always have a function main `func main`.*

---

3. Import fmt means that we are importing another package written by someone else to our main package . Fmt is a default package / standard library that comes along with our go download . It helps in printing input output. It is also known as format. We can see other standard library from https://pkg.go.dev/std

---

4. func is nothing but a way to represent function in golang. 

```go
func main (){

}
```
- func - Tells go that we are declaring a function
- main - sets the name of the function, we can use other names as well. 
- () - List of arguments to pass through a function
- {} - Function body, calling the function runs the code

---

5. We will always orgainze our go programs like this :
    
    a. package declaration 
    b. import other packages that we need 
    c. declare function , tell go to do the things 

    ```go 
    package main // a

    import "fmt" // b
 
    func main (){ // c
        fmt.Println("Hello")
    }
    ```
