package handlers

import "github.com/bwmarrin/discordgo"

func HandleStatus(bot*discordgo.Session) {
    status := discordgo.UpdateStatusData{
        Status: "dnd",
        Activities: []*discordgo.Activity{
            &discordgo.Activity{
                Name: "DiscordGo",
                Type: discordgo.ActivityTypeGame,
            },
        },
    }
    
    bot.UpdateStatusComplex(status)
}