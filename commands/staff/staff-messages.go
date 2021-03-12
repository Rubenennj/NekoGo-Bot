package staff 

import (
    "github.com/bwmarrin/discordgo"
    "../../functions"
    "strings"
    "strconv"
    "../../structures"
)

var Success = "Staff Commands Loaded"

var StaffMessagesCommand = structures.Command{
    Name: "staff-messages",
    Category: "staff",
    Desc: "Show staff messages",
    Aliases: []string{
        "sms",
    },
    Staff: true,
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        //depending on the user input 
        m := &discordgo.Member{}
        
        if len(args) == 0 {
            member, err := functions.GetMember(bot, msg.GuildID, msg.Author.ID)
            
            if err != nil {
                bot.ChannelMessageSend(msg.ChannelID, err.Error())
                return
            }
            
            m = member 
        } else {
            member, err := functions.FindMember(bot, msg.GuildID, args[0])
            
            if err != nil {
                bot.ChannelMessageSend(msg.ChannelID, err.Error())
                return
            }
            
            m = member
        }
        
        if functions.MemberIsStaff(m) == false {
            bot.ChannelMessageSend(msg.ChannelID, m.User.Username + " is not a staff member.")
            return
        }
        
        highest, unknownErr := functions.MemberHighestRole(bot, msg.GuildID, m)
        
        if unknownErr != nil {
            bot.ChannelMessageSend(msg.ChannelID, unknownErr.Error())
            return
        }
        
        data, err := functions.DatabaseMustGet("data_" + m.User.ID) 
        
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        channels := []string{}
        
        self := "You"
        
        if msg.Author.ID != m.User.ID {
            self = "They"
        }
        
        embed := &discordgo.MessageEmbed{
            Footer: &discordgo.MessageEmbedFooter{
                Text: "Made with DiscordGo",
                IconURL: bot.State.User.AvatarURL("4096"),
            },
            Color: highest.Color, 
            Title: m.User.Username + "'s Messages",
            Thumbnail: &discordgo.MessageEmbedThumbnail{
                URL: m.User.AvatarURL("4096"),
            },
            Description: self + " have sent a total of `" + strconv.FormatInt(int64(data.MessageCount), 10) + "` " + functions.Plural("message", int64(data.MessageCount)) + " in this guild.",
        }
        
        fields := []*discordgo.MessageEmbedField{}
        
        if len(data.Messages) > 0 {
            for _, id := range functions.SortMapStringInt(data.Messages) {
                count := data.Messages[id]
                
                ch, err := functions.GetChannel(bot, id)
                
                if err != nil {
                    continue
                }
                
                channels = append(channels, "[#" + ch.Name + "]: " + strconv.FormatInt(int64(count), 10) + " " + functions.Plural("message", int64(count)))
                
                if len(channels) == 15 {
                    //add a fancy text saying rest channels 
                    if len(data.Messages) > 15 {
                        channels = append(channels, "...and " + strconv.FormatInt(int64(len(data.Messages) - 15), 10) + " " + functions.Plural("channel", int64(len(data.Messages) - 15)) + " more")
                    }
                    
                    break
                }
            }
            
            fields = append(fields, &discordgo.MessageEmbedField{
                Name: "Channels " + self + " have chatted the most in:",
                Value: "```\n" + strings.Join(channels, "\n") + "```",
            })
        }
        
        if len(fields) > 0 {
            embed.Fields = fields
        }
        
        bot.ChannelMessageSendEmbed(msg.ChannelID, embed)
    },
}

var ExecStaffMessagesCmd = StaffMessagesCommand.Register()