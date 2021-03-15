package cache 

import "github.com/bwmarrin/discordgo"

var Roles = map[string]*discordgo.Role{}

var Guilds = map[string]*discordgo.Guild{}

var Members = map[string]*discordgo.Member{}

var Channels = map[string]*discordgo.Channel{} 

var Users = map[string]*discordgo.User{} 

var Messages = map[string]*discordgo.MessageCreate{}

var Presences = map[string]*discordgo.PresenceUpdate{} 