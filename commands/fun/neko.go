package fun 

import (
    "../../structures"
    "github.com/bwmarrin/discordgo"
    "../../functions"
)

type Request map[string]interface{}

var NekoCommand = structures.Command{
    Name: "neko",
    Desc: "Returns a neko from nekos.life api",
    Category: "fun",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        
        json := Request{}
        
        err := functions.RequestJSON("http://nekos.life/api/v2/img/neko", &json)
        if err != nil {
            bot.ChannelMessageSend(msg.ChannelID, err.Error())
            return
        }
        
        highestRole, _ := functions.MemberHighestRole(bot, msg.GuildID, msg.Member)
        
        bot.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
            Title: "Here is a neko for " + msg.Author.Username + ":",
            Color: highestRole.Color,
            Image: &discordgo.MessageEmbedImage{
                URL: json["url"].(string),
            },
            Footer: &discordgo.MessageEmbedFooter{
                Text: "Made with DiscordGo",
                IconURL: bot.State.User.AvatarURL("4096"),
            },
        })
    },
}


var NekoCommandExec = NekoCommand.Register()