package commands

import (
    "github.com/bwmarrin/discordgo"
    "../functions"
    "../structures"
)

var Perms = structures.Command{
    Name: "test",
    Desc: "Test command",
    Permissions: []string{
        "ADMINISTRATOR",
    },
    Args: 1,
    Examples: []string{"@Ruben", "392609934744748032"},
    Usages: []string{"<user>"},
    OwnerOnly: true,
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        user, err := functions.FindUser(bot, args[0])
        
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        bot.ChannelMessageSend(msg.ChannelID, `Found: ` + user.Username)
    },
}

var permExec = Perms.Register()