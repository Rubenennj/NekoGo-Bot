package functions 

import "github.com/bwmarrin/discordgo"

func RequestGuildMembers (bot *discordgo.Session, guildID string, query string, limit int) ([]*discordgo.Member, error) {
    finish := make(chan bool)
    
    members := []*discordgo.Member{}
    
    //chunk event 
    handler := bot.AddHandler(func (bot *discordgo.Session, g *discordgo.GuildMembersChunk) {
        if g.GuildID != guildID {
            return
        }
        
        for _, m := range g.Members {
            members = append(members, m)
        }
        
        if g.ChunkIndex + 1 == g.ChunkCount {
            finish <- true 
        }
    })
    
    //no one wants presences bruh
    err := bot.RequestGuildMembers(guildID, query, limit, false)
    if err != nil {
        //an error
        //delete chunk handler 
        handler()
        return nil, err 
    }
    
    <-finish
    //delete chunk handler
    handler()
    //return members 
    return members, nil 
}