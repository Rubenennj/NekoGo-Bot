package staff

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    "../../functions"
    "strings"
    "../../cache"
)

var NicknameCommand = structures.Command{
    Name: "nickname",
    Aliases: []string{"nick"},
    Args: 2,
    Fields: []string{
        "member",
        "nickname",
    },
    Staff: true,
    Usages: []string{
        "<member> <nickname | reset>",
    },
    Examples: []string{
        "@Ruben you cant modify me L",
    },
    Desc: "Changes the nickname of a member. Will also remove Change Username role if executing in pending username change channel.",
    Category: "staff",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        member, err := functions.FindMember(bot, msg.GuildID, args[0])
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        nick := member.User.Username 
        
        if args[1] != "reset" {
            nick = strings.Join(args[1:], " ")
        }
        
        err = bot.GuildMemberNickname(msg.GuildID, member.User.ID, nick)
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return 
        }
        
        content := []string{} 
        
        if msg.ChannelID == "671454243840065586" {
            if err := bot.GuildMemberRoleRemove(msg.GuildID, member.User.ID, "671453937685233664"); err != nil {
                bot.ChannelMessageSend(msg.ChannelID, err.Error())
                return 
            }
            
            content = append(content, `Successfully removed Username Change role from them.`)
        }
        
        oldn := member.User.Username 
        if member.Nick != "" {
            oldn = member.Nick
        }
        
        member.Nick = nick 
        
        if member.Nick == member.User.Username {
            member.Nick = ""
        }
        
        cache.Members[msg.GuildID + member.User.ID] = member 
        
        content = append(content, "Changed `" + oldn + "` nickname to `" + nick + "`.")
        
        bot.ChannelMessageSend(msg.ChannelID, strings.Join(content, "\n"))
    },
}

var NicknameCommandExec = NicknameCommand.Register()