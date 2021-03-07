package handlers

import "github.com/bwmarrin/discordgo"
import "fmt"

func MessageCreate(bot*discordgo.Session, msg*discordgo.MessageCreate) {
    
    fmt.Println("Receiced message " + msg.Content + " from user " + msg.Author.Username)
}