package handlers 

import (
    "github.com/bwmarrin/discordgo"
    "../cache"
    "../utils"
    "../functions"
)

func CacheStaffPresence(bot*discordgo.Session, presence*discordgo.PresenceUpdate) {
    //must be a member of bdfd 
    member, exists := cache.Members[utils.GuildID + presence.User.ID]
    if exists == false {
        return
    }
    
    if functions.MemberIsStaff(member) == false {
        return 
    }
    
    cache.Presences[presence.User.ID] = presence 
}