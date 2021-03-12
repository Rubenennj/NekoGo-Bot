package handlers

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "../config"
    "../structures"
    "../functions"
    "regexp"
)

func CommandHandler (bot *discordgo.Session, msg *discordgo.MessageCreate) {
    
    if msg.Author.Bot == true {
        return
    }
    
    if strings.HasPrefix(msg.Content, config.Prefix) == false {
        //check for bot ping 
        reg := regexp.MustCompile(`^<@!?` +  bot.State.User.ID + `>$`)
        if reg.Match([]byte(msg.Content)) == true {
            BotPing(bot, msg)
        }
        return
    }
    
    cargs := strings.Split(strings.TrimSpace(msg.Content[len(config.Prefix):]), " ")
    
    cmd, args := strings.ToLower(cargs[0]), cargs[1:]
    
    command := structures.Commands[cmd]
    
    if command.Name == "" {
        //aliases check 
        for _, c := range structures.Commands {
            if len(c.Aliases) > 0 {
                if functions.Includes(c.Aliases, cmd) == true {
                    command = c 
                    break
                }
            }
        }
        
        if command.Name == "" {
            //nothing found 
            return 
        }
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

func BotPing (bot *discordgo.Session, msg *discordgo.MessageCreate) {
    bot.ChannelMessageSend(msg.ChannelID, msg.Author.Mention() + " hi! My prefix is `"+ config.Prefix +"`.")
}