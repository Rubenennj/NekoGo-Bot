package developer

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    "../../functions"
    "strconv"
    "time"
    "strings"
)

type Request map[string]interface{}

var Success = "Developer Commands Loaded"

var TestCommand = structures.Command{
    Name: "test",
    OwnerOnly: true,
    Args: 4,
    Desc: "a test command",
    Info: "Only mentions and IDs are allowed.",
    Fields: []string{
        "channel",
        "winners",
        "duration",
        "title",
    },
    Usages: []string{
        "<channel> <winners> <duration> <title>",
    },
    Examples: []string{
        "#giveaways 1 1m idk",
    },
    Category: "developer",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        
        channel, err := functions.GetChannel(bot, args[0])
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        winners, intErr := strconv.ParseInt(args[1], 10, 64)
        if intErr != nil {
            bot.ChannelMessageSend(msg.ChannelID, intErr.Error())
            return
        } 
        
        duration, durErr := time.ParseDuration(args[2])
        if durErr != nil {
            bot.ChannelMessageSend(msg.ChannelID, durErr.Error())
            return
        }
        
        ms := int(int64(duration) / int64(time.Millisecond))
        
        title := strings.Join(args[3:], " ")
        
        gw := &functions.Giveaway{
            UserID: msg.Author.ID,
            Title: title,
            Winners: int(winners),
            Duration: ms,
            StartedAt: int(time.Now().UnixNano() / int64(time.Millisecond)),
            EndsAt: int(time.Now().UnixNano() / int64(time.Millisecond) + int64(ms)),
            GuildID: msg.GuildID,
            ChannelID: channel.ID,
            Ended: false,
        }
        
        _, gwErr := gw.Start(bot)
        if gwErr != nil {
            bot.ChannelMessageSend(msg.ChannelID, gwErr.Error())
            return
        }
        
        bot.ChannelMessageSend(msg.ChannelID, `Giveaway Created`)
    },
}


var TestExec = TestCommand.Register()