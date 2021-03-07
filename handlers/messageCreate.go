package handlers

import "github.com/bwmarrin/discordgo"
import "strings"
import "../config"

func MessageCreate(bot*discordgo.Session, msg*discordgo.MessageCreate) {
    if msg.Author.Bot == true {
        return
    }
    
    if strings.HasPrefix(msg.Content, config.Prefix) == false {
        return
    }
    
    content := strings.Split(strings.TrimSpace(msg.Content[len(config.Prefix):]), " ")
    
    cmd, args := strings.ToLower(content[0]), content[1:]
    
    command := Commands[cmd]
    
    if command.Name == "" {
        bot.ChannelMessageSend(msg.ChannelID, "Command " + cmd + " not found.")
        return
    } else {
        command.Run(bot, msg, args)
    }
}