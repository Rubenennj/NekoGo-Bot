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
    
    checker, err := CommandCheck(bot, msg, args, &command, true)
    
    if err != nil {
        bot.ChannelMessageSend(msg.ChannelID, err.Error())
        return
    }
    
    if checker == false {
        return
    }
    
    command.Run(bot, msg, args)
}