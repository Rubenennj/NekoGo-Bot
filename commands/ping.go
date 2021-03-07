package commands

import "../structures"
import "github.com/bwmarrin/discordgo"

var Ping = structures.Command{
    Name: "ping", 
    Description: "A ping command", 
    Run: func (bot*discordgo.Session, msg*discordgo.MessageCreate, args []string) {
        bot.ChannelMessageSend(msg.ChannelID, "Pong!")
    },
}