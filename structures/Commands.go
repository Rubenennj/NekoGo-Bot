package structures

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

var Commands = map[string]Command{}

type Command struct {
    Name string
    Desc string 
    Examples []string
    Permissions []string
    OwnerOnly bool 
    Args int
    Usages []string
    Run func (*discordgo.Session, *discordgo.MessageCreate, []string)
}

func (cmd Command) Register() *Command {
    //registers the command
    Commands[cmd.Name] = cmd 
    
    fmt.Println("Command " + cmd.Name + " loaded!")
    
    return &cmd
}