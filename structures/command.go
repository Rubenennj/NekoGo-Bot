package structures

import "github.com/bwmarrin/discordgo"

type Command struct {
    Name string
    Description string
    Run func (*discordgo.Session, *discordgo.MessageCreate, []string) 
}