package commands

import (
    "github.com/bwmarrin/discordgo"
    "../structures"
    "strconv"
)

var Ping = structures.Command{
    Name: "ping",
    Desc: "A ping command",
    OwnerOnly: false,
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        m, err := bot.ChannelMessageSend(msg.ChannelID, "Pinging...")
        
        if err != nil {
            return
        }
        
        time1, _ := m.Timestamp.Parse()
        time2, _ := msg.Timestamp.Parse()
        taken := strconv.FormatInt((time1.UnixNano() - time2.UnixNano()) / 1000000, 10)
        
        bot.ChannelMessageEdit(msg.ChannelID, m.ID, "Pong! " + taken + "ms")
    },
}

var exec = Ping.Register()