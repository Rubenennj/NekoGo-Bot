package handlers

import "github.com/bwmarrin/discordgo"
import "fmt"

type Command struct {
    Name string
    Description string
    Run func (*discordgo.Session, *discordgo.MessageCreate, []string) 
}

func (c Command) Register() int {
    Commands[c.Name] = c
    fmt.Println("Command " +  c.Name + " loaded!")
    
    return 1
}