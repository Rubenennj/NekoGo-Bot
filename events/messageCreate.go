package events

import (
    "github.com/bwmarrin/discordgo"
    "../handlers"
)

func MessageCreate (bot *discordgo.Session, msg *discordgo.MessageCreate) {
    handlers.CommandHandler(bot, msg)
}