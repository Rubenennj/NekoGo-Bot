package events 

import (
    "github.com/bwmarrin/discordgo"
    "../handlers"
)

func MessageReactionAdd (bot *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
    handlers.CollectorHandler(bot, reaction)
}