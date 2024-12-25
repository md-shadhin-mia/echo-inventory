package main

import (
    "fmt"
    "log"
)

func main() {
    // Using fmt to print a message
    fmt.Println("This is a message to the user.")

    // Using log to log an error
    log.Println("This is a log message.")
    
    // Logging a fatal error
    // log.Fatal("This is a fatal error message.") // This will exit the program
}