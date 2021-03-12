package functions 

import (
    "github.com/bwmarrin/discordgo"
    "time"
)

type ReactionCollector struct {
    MessageID string
    Max int 
    Time int 
    Filter func (*discordgo.MessageReactionAdd) bool 
    OnAdd func (*discordgo.MessageReactionAdd)
    OnEnd func () 
    OnRemove func (*discordgo.MessageReactionRemove)
}

var ReactionCollectors = map[string]*ReactionCollector{} 

func (c ReactionCollector) Stop() bool {
    delete(ReactionCollectors, c.MessageID)
    c.OnEnd()
    return true 
}

func (c ReactionCollector) Start() {
    ReactionCollectors[c.MessageID] = &c
    time.Sleep(time.Duration(c.Time)*time.Second)
    _, e := ReactionCollectors[c.MessageID]
    if e == true {
        c.OnEnd()
    }
    delete(ReactionCollectors, c.MessageID)
}
