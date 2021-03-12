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
        m, _ := bot.ChannelMessageSend(msg.ChannelID, "React to this message")
        
        collector := &functions.ReactionCollector{
            MessageID: m.ID,
            Time: 60,
            Filter: func (r*discordgo.MessageReactionAdd) bool {
                return msg.Author.ID == r.UserID
            },
            OnAdd: func (r*discordgo.MessageReactionAdd) {
                bot.ChannelMessageSend(msg.ChannelID, "You reacted with " + r.Emoji.MessageFormat())
            },
        }
        
        collector.Start()
    },
}


var TestExec = TestCommand.Register()