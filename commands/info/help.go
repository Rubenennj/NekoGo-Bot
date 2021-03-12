package info 

import (
    "github.com/bwmarrin/discordgo"
    "../../structures"
    "strings"
    "strconv"
    "../../config"
    "../../handlers"
)

var HelpCommand = structures.Command{
    Name: "help",
    Aliases: []string{"commands"},
    Desc: "List all available commands for you or get specific info on one",
    Category: "info",
    Run: func (bot *discordgo.Session, msg *discordgo.MessageCreate, args []string) {
        categories := map[string][]string{}
        
        embed := discordgo.MessageEmbed{
            Color: 9091583,
            Title: bot.State.User.Username + " Command List",
            Thumbnail: &discordgo.MessageEmbedThumbnail{
                URL: bot.State.User.AvatarURL("4096"),
            },
            Footer: &discordgo.MessageEmbedFooter{
                Text: "Requested by " + msg.Author.Username + ". | Made with DiscordGo",
                IconURL: msg.Author.AvatarURL("4096"),
            },
        }
        
        fields := []*discordgo.MessageEmbedField{}
        
        for _, cmd := range structures.Commands {
            checker, _ := handlers.CommandCheck(bot, msg, []string{}, &cmd, false)
            if checker == true {
                categories[cmd.Category] = append(categories[cmd.Category], "`"+ config.Prefix + cmd.Name +"`")
            }
        }
        
        for name, cmds := range categories {
            c := strconv.FormatInt(int64(len(cmds)), 10)
            fields = append(fields, &discordgo.MessageEmbedField{
                Name: "**__" + strings.ToUpper(name) + " ["+c+"]__**",
                Value: strings.Join(cmds, ", "),
            })
        }
        
        embed.Fields = fields 
        
        bot.ChannelMessageSendEmbed(msg.ChannelID, &embed)
    },
}

var HelpCommandExec = HelpCommand.Register()