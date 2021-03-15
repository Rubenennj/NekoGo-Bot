package functions

import (
    "github.com/bwmarrin/discordgo"
    "../utils"
    "errors"
    "../cache"
)

func FindMember (bot *discordgo.Session, guildID string, arg string) (*discordgo.Member, error) {
    byted := []byte(arg)
    
    if utils.UserMention.Match(byted) == true {
        id := utils.Symbols.ReplaceAllString(arg, "")
        
        member, err := GetMember(bot, guildID, id)
        
        if err != nil {
            return nil, err
        } else {
            return member, nil
        }
    } else if utils.UserID.Match(byted) == true {
        member, err := GetMember(bot, guildID, arg)
        
        if err != nil {
            return nil, err
        } else {
            return member, nil
        }
    } else {
        members, err := RequestGuildMembers(bot, guildID, arg, 1)
        if err != nil {
            return nil, err 
        }
        if len(members) == 0 {
            return nil, errors.New("Could not find any member with given query")
        }
        
        member := members[0]
        cache.Members[member.User.ID] = member 
        
        return member, nil 
    }
}