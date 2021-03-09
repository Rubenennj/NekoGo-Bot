package commands

import (
    "github.com/bwmarrin/discordgo"
    "../functions"
    "../structures"
    "strings"
)

var Perms = structures.Command{
    Name: "perms",
    Desc: "Return user perms",
    OwnerOnly: false,
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        perms, err := functions.MemberPermissions(bot, msg.GuildID, msg.Author.ID)
        
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        bot.ChannelMessageSend(msg.ChannelID, msg.Author.Username + " Permissions:\n" + strings.Join(functions.Goof(perms), ", "))
    },
}

var permExec = Perms.Register()