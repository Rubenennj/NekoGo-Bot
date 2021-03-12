package handlers 

import (
    "github.com/bwmarrin/discordgo"
    "../functions"
)

func CollectorHandler (bot *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
    data, exists := functions.ReactionCollectors[reaction.MessageID]
    if exists == false {
        return 
    }
    
    if data.Filter(reaction) == false {
        return 
    }
    
    data.OnAdd(reaction)
}