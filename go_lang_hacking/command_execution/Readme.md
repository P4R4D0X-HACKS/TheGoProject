# Go Command Executor

This Go package provides a simple way to execute system commands with multiple arguments. The main function demonstrates how to use the `executeCommand` function to run a command.

## Features

- Execute system commands with arguments.
- Redirect output and error to standard output and error.
- Simple and easy-to-use interface.

## Installation

Make sure you have Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

Clone this repository to your local machine:

```bash
git clone https://github.com/P4R4D0X-HACKS/TheGoProject.git
cd TheGoProject
```

## Usage
-----

The main function demonstrates how to use the `executeCommand` function.

### Example

```go
package main

import (
    "log"
    "os"
    "os/exec"
)

func executeCommand(command string, args_arr []string) (err error) {
    args := args_arr
    cmd := exec.Command(command, args...) // Create a command with multiple arguments
    cmd.Stdout = os.Stdout // Link standard output
    cmd.Stderr = os.Stderr // Link standard error

    err = cmd.Run() // Run the command
    if err != nil {
        log.Fatal("cmd.Run() failed ", err)
        return
    }
    return
}

func main() {
    command := "ls"
    executeCommand(command, []string{"-l"})
}
```

### Running the Example

To run the example, navigate to the directory containing your Go file and execute:

```bash
go run main.go
```

## Function Overview
-----------------

### `executeCommand(command string, args_arr []string) error`

-   **Parameters:**
    -   `command`: The command to be executed.
    -   `args_arr`: A slice of strings representing the arguments for the command.
-   **Returns:**
    -   An error if the command execution fails.

### Output

The output of the executed command will be printed to the console.