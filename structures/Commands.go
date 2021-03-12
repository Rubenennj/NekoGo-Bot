package structures

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

var Commands = map[string]Command{}

type Command struct {
    Name string
    Category string
    Desc string 
    Aliases []string 
    Examples []string
    Info string
    Staff bool 
    Permissions []string
    Fields []string
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