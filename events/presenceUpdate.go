package events 

import "github.com/bwmarrin/discordgo"
import "../handlers"

func PresenceUpdate (bot*discordgo.Session, presence*discordgo.PresenceUpdate) {
    handlers.CacheStaffPresence(bot, presence)
}