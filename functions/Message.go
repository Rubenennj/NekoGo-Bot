package functions 

import (
    "github.com/bwmarrin/discordgo"
    "time"
    "errors"
)

type ReactionCollectorData struct {
    MessageID string
    Time int 
    Ended bool 
    Filter func (*discordgo.MessageReactionAdd) bool
    OnCollect func (*discordgo.MessageReactionAdd) error 
    OnEnd func ()
    Handler func() 
}

func (c ReactionCollectorData) End() {
    if c.Ended == true {
        return 
    }
    c.Ended = true 
    c.Handler()
    c.OnEnd() 
}

func ReactionCollector(bot *discordgo.Session, data *ReactionCollectorData) {
  
    data.Handler = bot.AddHandler(func (bot *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
        if data.MessageID != reaction.MessageID {
            return 
        }
        
        if data.Filter(reaction) == false {
            return 
        }
        
        err := data.OnCollect(reaction)
        
        if err != nil {
            data.End() 
        }
    })
    
    time.AfterFunc(time.Duration(data.Time)*time.Second, func () {
        data.End() 
    })
}


func AwaitReaction(bot *discordgo.Session, t int, messageID string, filter func (*discordgo.MessageReactionAdd) bool) (*discordgo.MessageReactionAdd, error) {
    rr := &discordgo.MessageReactionAdd{}
    finish := make(chan bool)
    ended := false 
    
    handler := bot.AddHandler(func (bot *discordgo.Session, reaction*discordgo.MessageReactionAdd) {
        if messageID != reaction.MessageID {
            return 
        }
        
        if filter(reaction) == true {
            rr = reaction 
            ended = true 
            finish <- true 
        }
    })
    
    time.AfterFunc(time.Duration(t)*time.Second, func () {
        if ended == false {
            finish <- true 
        }
    })
    
    <- finish 
    //delete reaction handler
    handler()
    //no reactions collected
    if ended == false {
        return nil, errors.New("timeout")
    }
    return rr, nil 
}