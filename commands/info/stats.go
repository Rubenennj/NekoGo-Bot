package info 

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    //"../../functions"
    "time"
    "../../config"
    "runtime"
    "strconv"
)

var Success = "Info Commands Loaded"

var m runtime.MemStats 

var StatsCommand = structures.Command{
    Name: "stats",
    Desc: "show bot stats",
    Category: "info",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        runtime.ReadMemStats(&m)
        
        upt := time.Now().Sub(config.Uptime)
        
        embed := discordgo.MessageEmbed{
            Title: bot.State.User.Username + " Statistics",
            Color: 7055103,
            Description: "Library: DiscordGo\nCommands: " + strconv.FormatInt(int64(len(structures.Commands)), 10) + "\nGoroutines: " + strconv.FormatInt(int64(runtime.NumGoroutine()), 10) + "\nAlloc: " + ToMB(m.Alloc) + "mb\nTotalAlloc: " + ToMB(m.TotalAlloc) + "mb\nSys: " + ToMB(m.Sys) + "mb\nNumGC: " + strconv.FormatUint(uint64(m.NumGC), 10) + "\nUptime: " + upt.String(),
            Thumbnail: &discordgo.MessageEmbedThumbnail{
                URL: bot.State.User.AvatarURL("4096"),
            },
        }
        
        bot.ChannelMessageSendEmbed(msg.ChannelID, &embed)
    },
}

var StatsCmd = StatsCommand.Register()

func ToMB (i uint64) string {
    return strconv.FormatUint(i / 1024 / 1024, 10)
}