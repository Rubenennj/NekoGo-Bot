package structures

import (
    "fmt"
    "../handlers"
)

func RegisterCommand (cmd*Command) *Command {
    fmt.Println("Added " + cmd.Name + " to the handler")
    
    handlers.Commands[cmd.Name] = cmd
    
    return &cmd
}