package commands

import "../handlers"
import "github.com/bwmarrin/discordgo"
import "strconv"

var Success = "All bot commands loaded!"

var ping = handlers.Command{
    Name: "ping", 
    Description: "A ping command", 
    Run: func (bot*discordgo.Session, msg*discordgo.MessageCreate, args []string) {
        m, err := bot.ChannelMessageSend(msg.ChannelID, "Pinging...")
        
        if err != nil {
            return
        }
        
        time1, _ := msg.Timestamp.Parse()
        time2, _ := m.Timestamp.Parse()
        
        time := (time2.UnixNano() - time1.UnixNano()) / 1000000
        
        bot.ChannelMessageEdit(msg.ChannelID, m.ID, "Pong! " + strconv.FormatInt(int64(time), 10) + "ms")
    },
}

var execution = ping.Register()