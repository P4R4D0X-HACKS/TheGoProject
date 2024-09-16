# Bot Greetings in Go

This Go program demonstrates the use of interfaces to handle different types of bots that provide greetings in various languages. The program defines a `bot` interface and two types of bots: `englishBot` and `spanishBot`. Each bot type implements the `getGreeting` method to return a greeting in its respective language.

## Code Explanation

### Interface Definition

In Go, an interface is a type that specifies a set of method signatures. Any type that implements these methods is said to satisfy the interface. In this program, the `bot` interface is defined with a single method:

```go
type bot interface {
    getGreeting() string
}
```

### Bot Types

Two types, `englishBot` and `spanishBot`, are defined to implement the `bot` interface. Each type has its own implementation of the `getGreeting` method:

```go
type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
    return "Hi There!!"
}

func (sb spanishBot) getGreeting() string {
    return "Hola"
}
```

### Main Function

In the `main` function, instances of `englishBot` and `spanishBot` are created. The `printGreeting` function is called with each bot instance. The `printGreeting` function takes a `bot` interface as a parameter, allowing it to accept any type that satisfies the `bot` interface:


```go
func main() {
    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}
```

# References in Go

In Go, interfaces are a powerful feature that allows you to define the behavior of objects without specifying their concrete types. This promotes flexibility and code reusability. Here are some key points about interfaces in Go:

- **An interface type specifies a set of method signatures.**
- **A type satisfies an interface by implementing its methods.**
- **Interfaces enable polymorphism**, allowing you to write functions that can operate on different types that satisfy the same interface.
