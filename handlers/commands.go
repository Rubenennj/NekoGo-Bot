package handlers

//import "github.com/bwmarrin/discordgo"
import "../commands"
import "../structures"

var Commands = map[string]structures.Command{
    "ping": commands.Ping,
}
