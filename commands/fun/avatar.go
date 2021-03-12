package fun 

import (
    "github.com/bwmarrin/discordgo"
    "../../structures"
    "../../functions"
)

var Success = "Fun Commands Loaded"

var AvatarCommand = structures.Command{
    Name: "avatar",
    Aliases: []string{"av"},
    Desc: "Returns yours or someone else's avatar",
    Category: "fun",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        user := &discordgo.User{}
        
        if len(args) > 0 {
            u, err := functions.FindUser(bot, args[0])
            
            if err != nil {
                bot.ChannelMessageSend(msg.ChannelID, err.Error())
                return
            }
            
            user = u 
        } else {
            user = msg.Author 
        }
        
        embed := &discordgo.MessageEmbed{
            Title: "Here is " + user.Username + "#" + user.Discriminator + " avatar:",
            Image: &discordgo.MessageEmbedImage{
                URL: user.AvatarURL("4096"),
            },
            Footer: &discordgo.MessageEmbedFooter{
                IconURL: msg.Author.AvatarURL("4096"),
                Text: "Requested by " + msg.Author.Username + " | Made with DiscordGo",
            },
        }
        
        member, err := functions.GetMember(bot, msg.GuildID, user.ID)
        
        if err == nil {
            role, e := functions.MemberHighestRole(bot, msg.GuildID, member)
            
            if e == nil {
                embed.Color = role.Color 
            }
        }
        
        bot.ChannelMessageSendEmbed(msg.ChannelID, embed) 
    },
}

var AvatarCommandExec = AvatarCommand.Register()