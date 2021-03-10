package functions

import (
    "github.com/bwmarrin/discordgo"
    "../utils"
    "errors"
)

func FindUser (bot *discordgo.Session, arg string) (*discordgo.User, error) {
    byted := []byte(arg)
    
    if utils.UserMention.Match(byted) == true {
        id := utils.Symbols.ReplaceAllString(arg, "")
        
        user, err := bot.User(id)
        
        if err != nil {
            return nil, err
        } else {
            return user, nil
        }
    } else if utils.UserID.Match(byted) == true {
        user, err := bot.User(arg)
        
        if err != nil {
            return nil, err
        } else {
            return user, nil
        }
    } else {
        return nil, errors.New("???")
    }
}