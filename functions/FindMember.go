package functions

import (
    "github.com/bwmarrin/discordgo"
    "../utils"
    "errors"
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
        return nil, errors.New("Query out of Regex")
    }
}