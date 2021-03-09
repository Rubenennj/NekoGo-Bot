package handlers

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "../config"
    "../structures"
)

func CommandHandler (bot *discordgo.Session, msg *discordgo.MessageCreate) {
    
    if msg.Author.Bot == true {
        return
    }
    
    if strings.HasPrefix(msg.Content, config.Prefix) == false {
        return
    }
    
    cargs := strings.Split(strings.TrimSpace(msg.Content[len(config.Prefix):]), " ")
    
    cmd, args := strings.ToLower(cargs[0]), cargs[1:]
    
    command := structures.Commands[cmd]
    
    if command.Name == "" {
        return
    }
    
    if msg.GuildID == "" {
        return
    }
    
    command.Run(bot, msg, args)
}