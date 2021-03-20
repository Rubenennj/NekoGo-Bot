package functions 

import (
    "github.com/bwmarrin/discordgo"
    "encoding/json"
    "../config"
    "time"
    "github.com/hako/durafmt"
    "strings"
    "math/rand"
    "fmt"
    "errors"
    "strconv"
)

type Giveaway struct {
    Title string `json:title`
    Winners int `json:winners`
    Duration int `json:duration`
    EndsAt int `json:endsAt`
    StartedAt int `json:StartedAt`
    Ended bool `json:ended`
    MessageID string `json:messageID`
    ChannelID string `json:channelID`
    GuildID string `json:guildID`
    UserID string `json:userID`
}

func FetchGiveaways(bot *discordgo.Session) {
    config.Database.Scan([]byte("giveaway_"), func (key []byte) error {
        data, _ := config.Database.Get(key)
        gw := Giveaway{}
        err := json.Unmarshal(data, &gw)
        if err != nil {
            fmt.Println(err)
            return err 
        }
        
        if gw.Ended == false {
            fmt.Println("Fetched giveaway " + gw.MessageID + " (" + gw.Title + ") on " + strconv.FormatBool(gw.Ended) + " state.")
            
            gw.Update(bot) 
        }
        
        return nil 
    })
}

func (gw Giveaway) Embed(bot *discordgo.Session) (*discordgo.MessageEmbed, error) {
    msRemaining := int64(gw.EndsAt) - (time.Now().UnixNano() / int64(time.Millisecond))
    dur, err := durafmt.ParseString(time.Duration(msRemaining * int64(time.Millisecond)).String())
    if err != nil {
        return nil, err
    }
    
    parsedDur := []string{}
    
    for _, val := range strings.Split(dur.String(), " ") {
        if Includes([]string{
            "microseconds",
            "milliseconds",
        }, val) {
            parsedDur = parsedDur[:len(parsedDur)-1]
            break
        } else {
            parsedDur = append(parsedDur, val)
        }
    }
    
    guild, gErr := GetGuild(bot, gw.GuildID)
    if gErr != nil {
        return nil, gErr
    }
    
    embed := &discordgo.MessageEmbed{
        Color: 8038911,
        Author: &discordgo.MessageEmbedAuthor{
            Name: "🎉 BDFD GIVEAWAY 🎉",
        },
        Title: gw.Title,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: guild.IconURL(),
        },
        Description: "**Hoster**: <@" + gw.UserID + ">\n**Time Remaining**: " + strings.Join(parsedDur, " "),
        Footer: &discordgo.MessageEmbedFooter{
            Text: "React with 🎉 to enter the giveaway! | Go power",
        },
    }
    
    return embed, nil 
}

func (gw *Giveaway) Start(bot *discordgo.Session) (string, error) {
    embed, embedErr := gw.Embed(bot)
    if embedErr != nil {
        return "", embedErr
    }
    
    msg, msgErr := bot.ChannelMessageSendEmbed(gw.ChannelID, embed)
    if msgErr != nil {
        return "", msgErr
    }
    
    bot.MessageReactionAdd(msg.ChannelID, msg.ID, "🎉")
    
    gw.MessageID = msg.ID 
    
    s, e := gw.Save()
    if e != nil {
        return "", e
    }
    
    time.AfterFunc(time.Duration(1), func() {
        gw.Update(bot) 
    })
    
    return s, nil 
}

func (gw Giveaway) Random() int64 {
    //seed for minecraft :v
    rand.Seed(time.Now().UnixNano())
    
    rn := rand.Int63n(int64(120000)) + int64(120000)
    
    if int64(gw.EndsAt) - time.Now().UnixNano() / int64(time.Millisecond) < 1 {
        return 1
    }
    
    if (int64(gw.EndsAt) - time.Now().UnixNano() / int64(time.Millisecond)) - rn < 1 {
        rn = int64(gw.EndsAt) - time.Now().UnixNano() / int64(time.Millisecond)
    }
    
    return rn 
}

func (gw Giveaway) Save() (string, error) {
    if gw.MessageID == "" {
        return "", errors.New("Giveaway without Giveaway.MessageID string type field provided")
    }
    
    data, err := json.Marshal(gw)
    if err != nil {
        return "", err 
    }
    
    err = config.Database.Put([]byte("giveaway_" + gw.MessageID), data)
    if err != nil {
        return "", err 
    }
    
    return string(data), nil 
}

func (gw *Giveaway) End(bot *discordgo.Session) {
    gw.Ended = true 
    users := gw.FetchReactions(bot)
    winners := gw.GetWinners(users)
    if len(winners) == 0 {
        bot.ChannelMessageSend(gw.MessageID, "No one won the giveaway L")
        return 
    } 
    
    plr := Plural("Winner", int64(len(winners)))
    
    guild, _ := GetGuild(bot, gw.GuildID)
    
    embed := &discordgo.MessageEmbed{
        Color: 8038911,
        Author: &discordgo.MessageEmbedAuthor{
            Name: "🎉 BDFD GIVEAWAY (ENDED) 🎉",
        },
        Title: gw.Title,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: guild.IconURL(),
        },
        Description: "**Hoster**: <@" + gw.UserID + ">\n**" + plr + "**: " + strings.Join(winners, ", "),
        Footer: &discordgo.MessageEmbedFooter{
            Text: "Thanks to those who participated! | Go power",
        },
    }
    
    bot.ChannelMessageEditEmbed(gw.ChannelID, gw.MessageID, embed)
    
    self := "Congratulations " + winners[0] + "! You won **" + gw.Title + "**."
    
    if len(winners) > 1 {
        self = "Congratulations " + strings.Join(winners, ", ") + "! You won **" + gw.Title + "**."
    }
    
    gw.Save()
    
    bot.ChannelMessageSend(gw.ChannelID, self)
}

func (gw Giveaway) GetWinners(users []*discordgo.User) []string {
    if len(users) == 0 {
        return []string{}
    }
    
    winners := []string{}
    
    for i := 0; i < gw.Winners; i++ {
        rand.Seed(time.Now().UnixNano())
        
        index := rand.Int63n(int64(len(users))) 
        
        winners = append(winners, users[index].Mention())
    }
    
    return winners
}

func (gw Giveaway) Update(bot *discordgo.Session) {
    wait := gw.Random()
    
    if gw.Ended {
        return 
    }
    
    time.Sleep(time.Duration(wait * int64(time.Millisecond)))
    
    if int64(gw.EndsAt) - time.Now().UnixNano() / int64(time.Millisecond) < 10 {
        gw.End(bot)
        return
    }
    
    embed, err := gw.Embed(bot)
    if err != nil {
        fmt.Println(err.Error())
        return 
    }
    
    _, mErr := bot.ChannelMessageEditEmbed(gw.ChannelID, gw.MessageID, embed)
    if mErr != nil {
        fmt.Println(mErr.Error())
    }
    
    gw.Update(bot)
}

func (gw Giveaway) FetchReactions(bot *discordgo.Session) []*discordgo.User {
    lastID := "1"
    
    collection := []*discordgo.User{}
    
    y := 0 
    
    for true {
        if y >= 100 {
            return collection
        }
        
        users, _ := bot.MessageReactions(gw.ChannelID, gw.MessageID, "🎉", 100, "", lastID)
        
        if len(users) == 1 {
            return collection
        }
        
        for _, user := range users {
            if user.Bot {
                continue
            }
            
            collection = append(collection, user)
            
            lastID = user.ID
        }
        
        y++ 
    }
    
    return collection 
}