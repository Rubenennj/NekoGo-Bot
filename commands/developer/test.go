package developer

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    "../../functions"
    "../../cache"
    "strconv"
    "errors"
    "strings"
)

var Success = "Developer Commands Loaded"

var TestCommand = structures.Command{
    Name: "test",
    OwnerOnly: true,
    Args: 1,
    Desc: "a test command",
    Info: "Only mentions and IDs are allowed.",
    Fields: []string{
        "user",
    },
    Usages: []string{
        "<user>",
    },
    Examples: []string{
        "@Ruben",
        "356950275044671499",
    },
    Category: "developer",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        y := 0
        
        if len(cache.Presences) == 0 {
            bot.ChannelMessageSend(msg.ChannelID, ":x: No presences to retrieve!")
            return 
        }
        
        m, _ := bot.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
            Title: "React to start",
        })
         
        presences := []*discordgo.PresenceUpdate{}
        
        for _, presence := range cache.Presences {
            presences = append(presences, presence)
        }
        
        functions.ReactionCollector(bot, &functions.ReactionCollectorData{
            Time: 60,
            MessageID: m.ID,
            Filter: func (r *discordgo.MessageReactionAdd) bool {
                return msg.Author.ID == r.UserID && functions.Includes([]string{"◀️", "▶️", "❌"}, r.Emoji.Name)
            },
            OnEnd: func () {
                bot.ChannelMessageDelete(msg.ChannelID, m.ID)
            },
            OnCollect: func (r *discordgo.MessageReactionAdd) error {
                if r.Emoji.Name == "❌" {
                    return errors.New("close")
                }
                
                bot.MessageReactionRemove(msg.ChannelID, m.ID, r.Emoji.Name, msg.Author.ID)
                
                if r.Emoji.Name == "◀️" {
                    if y - 1 < 0 {
                        y = len(presences) - 1 
                    } else {
                        y--
                    }
                } else {
                    if y + 1 >= len(presences) {
                        y = 0 
                    } else {
                        y++ 
                    }
                }
                
                presence := presences[y]
                
                embed := &discordgo.MessageEmbed{
                    Color: 13428479,
                    Author: &discordgo.MessageEmbedAuthor{
                        Name:"Presence List (Test)",
                    },
                    Footer: &discordgo.MessageEmbedFooter{
                        IconURL: msg.Author.AvatarURL("4096"),
                        Text: "Page "+ strconv.FormatInt(int64(y+1), 10) +" / "+ strconv.FormatInt(int64(len(presences)), 10) +" | Requested by " + msg.Author.Username + " | Made with DiscordGo", 
                    },
                }
                
                member, err := functions.GetMember(bot, msg.GuildID, presence.User.ID)
                if err != nil {
                    embed.Description = ":x: Failed to retrieve member " + presence.User.Mention()
                }
                
                if err == nil {
                    embed.Title = member.User.Username + "#" + member.User.Discriminator 
                    
                    embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
                        URL: member.User.AvatarURL("4096"),
                    }
                    
                    fields := []*discordgo.MessageEmbedField{
                        {
                            Name: "Status",
                            Value: string(presence.Status),
                        },
                    }
                    
                    if len(presence.Activities) > 0 {
                        content := []string{}
                        for _, activity := range presence.Activities {
                            content = append(content, activity.Name)
                        }
                        fields = append(fields, &discordgo.MessageEmbedField{
                            Name: "Activities",
                            Value: strings.Join(content, "\n"),
                        })
                    }
                    
                    embed.Fields = fields 
                }
                
                bot.ChannelMessageEditEmbed(msg.ChannelID, m.ID, embed)
                
                return nil 
            },
        })
        
        bot.MessageReactionAdd(msg.ChannelID, m.ID, "◀️")
        bot.MessageReactionAdd(msg.ChannelID, m.ID, "❌")
        bot.MessageReactionAdd(msg.ChannelID, m.ID, "▶️")
    },
}


var TestExec = TestCommand.Register()