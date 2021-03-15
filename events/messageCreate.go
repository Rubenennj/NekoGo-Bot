package events

import (
    "github.com/bwmarrin/discordgo"
    "../handlers"
    "../functions"
    "../utils"
    "fmt"
    "../cache"
)

func MessageCreate (bot *discordgo.Session, msg *discordgo.MessageCreate) {
    handlers.CommandHandler(bot, msg)
    
    if msg.GuildID == utils.GuildID {
        cache.Members[msg.Author.ID] = msg.Member 
        
        member, err := functions.GetMember(bot, utils.GuildID, msg.Author.ID)
        if err != nil {
            return
        }
        if functions.ArrayMapIncludes(member.Roles, utils.StaffRoles) == false {
            return 
        }
        
        err = functions.AddMessagesWithData("data_" + msg.Author.ID, 1, msg)
        
        if err != nil {
            fmt.Println(err.Error())
        }
    }
}
