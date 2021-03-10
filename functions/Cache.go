package functions

import "github.com/bwmarrin/discordgo"
import "../cache"

func GetChannel (bot *discordgo.Session, channelID string) (*discordgo.Channel, error) {
    ch, exists := cache.Channels[channelID]
    
    if exists != true {
        channel, err := bot.Channel(channelID)
        if err != nil {
            return nil, err
        }
        
        cache.Channels[channel.ID] = channel
        
        return channel, nil 
    }
    
    return ch, nil
}

func GetMember (bot *discordgo.Session, guildID string, memberID string) (*discordgo.Member, error) {
    memb, exists := cache.Members[memberID]
    
    if exists != true {
        member, err := bot.GuildMember(guildID, memberID)
        if err != nil {
            return nil, err 
        }
        
        cache.Members[memberID] = member 
        
        return member, nil
    }
    
    return memb, nil
}

func GetGuild (bot *discordgo.Session, guildID string) (*discordgo.Guild, error) {
    g, exists := cache.Guilds[guildID]
    
    if exists != true {
        guild, err := bot.Guild(guildID)
        if err != nil {
            return nil, err
        }
        
        cache.Guilds[guild.ID] = guild
        
        for _, role := range guild.Roles {
            cache.Roles[role.ID] = role
        }
        
        return guild, nil
    }
    
    return g, nil 
}