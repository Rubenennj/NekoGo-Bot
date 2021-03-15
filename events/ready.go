package events 

import(
    "../handlers"
    "github.com/bwmarrin/discordgo"
)

func Ready (bot *discordgo.Session, ready*discordgo.Ready) {
    handlers.HandleStatus(bot)
}