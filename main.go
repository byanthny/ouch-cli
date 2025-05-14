package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("ouch...")
	// fmt.Println("use exit to quit")
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("> ")
        
        if !scanner.Scan() {
            break
        }
        
        command := strings.TrimSpace(scanner.Text())
        
		switch {
			case command == "exit":
				fmt.Println("the end.")
				return
			case strings.HasPrefix(command, "-"):
				subCommand := strings.TrimPrefix(command, "-")
				switch subCommand {
					case "help":
					fmt.Println("lol. use exit to quit.")
					default:
					fmt.Printf("unknown command: %s\n", subCommand)
				}
			default:
				fmt.Println("use - for commands, e.g. -help")
		}

    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
}