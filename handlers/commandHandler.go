package handlers

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "../config"
    "../structures"
    "../functions"
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
    
    if command.OwnerOnly == true && msg.Author.ID != config.Owner {
        return 
    }
    
    if len(command.Permissions) > 0 {
        permissions, err := functions.MemberPermissions(bot, msg.GuildID, msg.Author.ID)
        
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        for _, perm := range command.Permissions {
            if functions.Has(permissions, perm) == false {
                bot.ChannelMessageSend(msg.ChannelID, ":x: " + msg.Author.Mention() +  " you're missing the `" + strings.Title(strings.ToLower(strings.Replace(perm, "_", " ", 1))) + "` permission to use this command.")
                return
            }
        }
    }
    
    if command.Args > 0 && len(args) < command.Args {
        fields := []*discordgo.MessageEmbedField{
            &discordgo.MessageEmbedField{
                Name: "Usage",
                Value: functions.Join(functions.Map(command.Usages, func (str string) string {
                    return "```\n" + config.Prefix + command.Name + " " + str + "```"
                }), "\n"),
            },
        }
        desc := ""
        if command.Desc != "" {
            desc = command.Desc
        }
        
        if len(command.Examples) > 0 {
            fields = append(fields, &discordgo.MessageEmbedField{
                Name: "Example",
                Value: "```\n" + functions.Join(functions.Map(command.Examples, func (str string) string {
                    return config.Prefix + command.Name + " " + str
                }), "\n") + "```",
            })
        }
        
        embed := &discordgo.MessageEmbed{
            Title: "Invalid Usage!",
            Thumbnail: &discordgo.MessageEmbedThumbnail{
                URL: bot.State.User.AvatarURL("4096"),
            },
            Description: desc,
            Fields: fields,
            Color: 16714254,
        }
        bot.ChannelMessageSendEmbed(msg.ChannelID, embed)
        return
    }
    
    command.Run(bot, msg, args)
}