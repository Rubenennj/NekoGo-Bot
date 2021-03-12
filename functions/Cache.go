package functions

import "github.com/bwmarrin/discordgo"
import "../cache"

func GetChannel (bot *discordgo.Session, channelID string) (*discordgo.Channel, error) {
    //check if we have this channel in oir own cache
    ch, exists := cache.Channels[channelID]
    
    if exists != true {
        //check for dgo cache
        channel, err := bot.State.Channel(channelID)
        if err != nil {
            //request api for the channel 
            channel, err = bot.Channel(channelID)
            
            if err != nil {
                //failed to get channel 
                return nil, err
            }
        }
        
        cache.Channels[channel.ID] = channel
        
        return channel, nil 
    }
    
    return ch, nil
}

func GetMember (bot *discordgo.Session, guildID string, memberID string) (*discordgo.Member, error) {
    //check for member in our own cache
    memb, exists := cache.Members[guildID + memberID]
    
    if exists != true {
        //check for dgo cache
        member, err := bot.State.Member(guildID, memberID)
        
        if err != nil {
            //call the api 
            member, err = bot.GuildMember(guildID, memberID)
            if err != nil {
                //member not found
                return nil, err 
            }
        }
        
        cache.Members[guildID + memberID] = member 
        
        return member, nil
    }
    
    return memb, nil
}

func GetUser (bot *discordgo.Session, userID string) (*discordgo.User, error) {
    //check for our own cache 
    u, exists := cache.Users[userID]
    
    if exists == false {
        //request api as dgo does not have user cache (?)
        user, err := bot.User(userID)
        if err != nil {
            return nil, err 
        }
        
        cache.Users[userID] = user 
        
        return user, nil
    }
    
    return u, nil 
}

func GetGuild (bot *discordgo.Session, guildID string) (*discordgo.Guild, error) {
    //check for our own cache
    g, exists := cache.Guilds[guildID]
    
    if exists != true {
        //check for dgo cache 
        guild, err := bot.State.Guild(guildID)
        if err != nil {
            //call the api 
            guild, err = bot.Guild(guildID)
            if err != nil {
                //failed to get guild 
                return nil, err 
            }
        }
        
        cache.Guilds[guild.ID] = guild
        
        for _, role := range guild.Roles {
            cache.Roles[role.ID] = role
        }
        
        return guild, nil
    }
    
    return g, nil 
}

func GetRole (bot *discordgo.Session, guildID string, roleID string) (*discordgo.Role, error) {
    //check our own cache
    r, exists := cache.Roles[roleID]
    
    if exists == false {
        //check dgo cache 
        role, err := bot.State.Role(guildID, roleID)
        
        if err != nil {
            //request all api roles
            roles, err2 := bot.GuildRoles(guildID)
            if err2 != nil {
                return nil, err2
            }
            
            for _, rol := range roles {
                cache.Roles[rol.ID] = rol 
                if rol.ID == roleID {
                    role = rol 
                }
            }
        }
        
        return role, nil
    }
    
    return r, nil 
}