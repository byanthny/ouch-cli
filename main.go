package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("ouch...")
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("_ ")
        
        if !scanner.Scan() {
            break
        }
        
        command := strings.TrimSpace(scanner.Text())
        
        if command == "exit" {
            fmt.Println("the end.")
            break
        }
        
        fmt.Printf("command: %s\n", command)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
}