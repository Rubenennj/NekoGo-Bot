package info 

import (
    "github.com/bwmarrin/discordgo"
    "../../structures"
    "strconv"
)

var PingCommand = structures.Command{
    Name: "ping",
    Category: "info",
    Desc: "Returns the bot latency",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        m, err := bot.ChannelMessageSend(msg.ChannelID, "Pinging...")
        if err != nil {
            return 
        }
        
        time1, _ := msg.Timestamp.Parse()
        time2, _ := m.Timestamp.Parse()
        
        time := (time2.UnixNano() - time1.UnixNano()) / 1000000
        
        bot.ChannelMessageEdit(m.ChannelID, m.ID, "Pong! " + strconv.FormatInt(int64(time), 10) + "ms")
    },
}

var PingCommandExec = PingCommand.Register()