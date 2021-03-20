package events 

import(
    "../handlers"
    "github.com/bwmarrin/discordgo"
    "../functions"
)

func Ready (bot *discordgo.Session, ready*discordgo.Ready) {
    handlers.HandleStatus(bot)
    functions.FetchGiveaways(bot)
}