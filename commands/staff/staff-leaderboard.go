package staff 

import (
    "github.com/bwmarrin/discordgo"
    "../../functions"
    "../../structures"
    "strings"
    "strconv"
)

var StaffLeaderboardCommand = structures.Command{
    Name: "staff-leaderboard",
    Desc: "Returns a leaderboard of staff messages",
    Category: "staff",
    Staff: true,
    Aliases: []string{
        "stafflb",
        "staff-lb",
        "slb",
        "sleaderboard",
    },
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        guild, _ := functions.GetGuild(bot, msg.GuildID) 
        
        bot.ChannelTyping(msg.ChannelID)
        
        all, err := functions.DatabaseGetKeysWith("data_")
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        keys := functions.DatabaseSortByCount(all)
        
        content := []string{}
        
        embed := discordgo.MessageEmbed{
            Color: 9091583,
            Title: "Staff Message Leaderboard",
            Thumbnail: &discordgo.MessageEmbedThumbnail{
                URL: guild.IconURL(),
            },
            Footer: &discordgo.MessageEmbedFooter{
                IconURL: msg.Author.AvatarURL("4096"),
                Text: "Requested by " + msg.Author.Username + ". | Made with DiscordGo",
            },
        }
        
        top := 1 
        
        for _, index := range keys {
            data := all[index]
            
            id := strings.Split(index, "_")[1]
            
            member, err := functions.GetMember(bot, msg.GuildID, id)
            
            if err != nil || functions.MemberIsStaff(member) == false || member.User.Bot == true {
                continue
            }
            
            count := strconv.FormatInt(int64(data.MessageCount), 10)
            plr := functions.Plural("message", int64(data.MessageCount))
            table := ""
            t := "`" + strconv.FormatInt(int64(top), 10) + "#` - "
            
            if member.Nick == "" {
                table = member.User.Username + " [" + member.Mention() + "]"
            } else {
                table = member.Nick + " [" + member.Mention() + "]"
            }
            
            content = append(content, t + table + ": " + count + " " + plr)
            
            top++ 
        }
        
        embed.Description = strings.Join(content, "\n")
        
        _, err = bot.ChannelMessageSendEmbed(msg.ChannelID, &embed)
        
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
    },
}

var StaffLeaderboardCommandExec = StaffLeaderboardCommand.Register()