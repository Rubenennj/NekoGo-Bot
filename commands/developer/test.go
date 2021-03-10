package developer

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    "../../functions"
)

var Success = "Developer Commands Loaded"

var TestCommand = structures.Command{
    Name: "test",
    OwnerOnly: true,
    Args: 1,
    Desc: "a test command",
    Info: "Only mentions and IDs are allowed.",
    Fields: []string{
        "user",
    },
    Usages: []string{
        "<user>",
    },
    Examples: []string{
        "@Ruben",
        "356950275044671499",
    },
    Category: "developer",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        user, err := functions.FindUser(bot, args[0])
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        bot.ChannelMessageSend(msg.ChannelID, "User: " + user.Username)
    },
}


var TestExec = TestCommand.Register()