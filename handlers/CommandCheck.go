package handlers

import (
    "github.com/bwmarrin/discordgo"
    "../structures"
    "../functions"
    "../config"
    "../utils"
    "time"
)

func Attach(bot *discordgo.Session, msg *discordgo.MessageCreate, messageID string) {
    err := bot.MessageReactionAdd(msg.ChannelID, messageID, "❌")
    if err != nil {
        return 
    }
    
    _, err = functions.AwaitReaction(bot, 15, messageID, func (r*discordgo.MessageReactionAdd) bool {
        return r.Emoji.Name == "❌" && r.UserID == msg.Author.ID 
    })
    
    if err != nil && err.Error() == "timeout" {
        bot.MessageReactionRemove(msg.ChannelID, messageID, "❌", bot.State.User.ID)
    } else {
        bot.ChannelMessageDelete(msg.ChannelID, messageID) 
    }
}

func CommandCheck (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string, command *structures.Command, sendMessage bool) (bool, error) {
    if command.OwnerOnly == true && config.Owner != msg.Author.ID {
        return false, nil
    }
    
    if command.Staff == true {
        member, err := functions.GetMember(bot, utils.GuildID, msg.Author.ID)
        if err != nil {
            return false, err 
        }
        
        if functions.ArrayMapIncludes(member.Roles, utils.StaffRoles) == false {
            if sendMessage == true {
                StaffHandler(bot, msg, command)
            }
            return false, nil 
        }
    }
    
    if sendMessage == true && command.Args > 0 && len(args) < command.Args {
        ArgsHandler(bot, msg, args, command)
        return false, nil
    }
    
    return true, nil
}

func ArgsHandler (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string, command *structures.Command) {
    embed := discordgo.MessageEmbed{
        Color: 16711680,
        Timestamp: time.Now().Format("2006-01-02T15:04:05.000Z"),
        Author: &discordgo.MessageEmbedAuthor{
            Name: msg.Author.Username + "#" + msg.Author.Discriminator,
            IconURL: msg.Author.AvatarURL("4096"),
        },
        Footer: &discordgo.MessageEmbedFooter{
            Text: "Tap to dismiss. | Made with DiscordGo",
        },
        Title: "Missing Arguments",
        Description: msg.Author.Mention() + " you did not give the `" + command.Fields[len(args)] + "` argument.",
        Thumbnail: &discordgo.MessageEmbedThumbnail{
        URL: bot.State.User.AvatarURL("4096"),
        },
    }
    
    fields := []*discordgo.MessageEmbedField{}
    
    if len(command.Usages) > 0 {
        fields = append(fields, &discordgo.MessageEmbedField{
            Name: "Usage",
            Value: "```\n" + functions.Join(functions.Map(command.Usages, func (str string) string {
                return config.Prefix + command.Name + " " + str
            }), "\n") + "```",
        })
    }
    
    if len(command.Examples) > 0 {
        fields = append(fields, &discordgo.MessageEmbedField{
            Name: "Example",
            Value: "```\n" + functions.Join(functions.Map(command.Examples, func (str string) string {
                return config.Prefix + command.Name + " " + str
            }), "\n") + "```",
        })
    }
    
    if len(fields) > 0 {
        embed.Fields = fields
    }
    
    m, err := bot.ChannelMessageSendEmbed(msg.ChannelID, &embed)
    if err != nil {
        return
    }
    
    Attach(bot, msg, m.ID)
}

func StaffHandler (bot *discordgo.Session, msg *discordgo.MessageCreate, command *structures.Command) {
    embed := discordgo.MessageEmbed{
        Color: 16711680,
        Timestamp: time.Now().Format("2006-01-02T15:04:05.000Z"),
        Author: &discordgo.MessageEmbedAuthor{
            Name: msg.Author.Username + "#" + msg.Author.Discriminator,
            IconURL: msg.Author.AvatarURL("4096"),
        },
        Footer: &discordgo.MessageEmbedFooter{
            Text: "Tap to dismiss. | Made with DiscordGo",
        },
        Title: "Missing Permissions",
        Description: msg.Author.Mention() + " You need to be a staff member to use `" + command.Name + "`.",
        Thumbnail: &discordgo.MessageEmbedThumbnail{
        URL: bot.State.User.AvatarURL("4096"),
        },
    }
    
    m, err := bot.ChannelMessageSendEmbed(msg.ChannelID, &embed)
    if err != nil {
        return 
    }
     
    Attach(bot, msg, m.ID)
}