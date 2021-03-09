package functions 

import "github.com/bwmarrin/discordgo"
import "errors"
import "../structures"

func MemberPermissions(bot *discordgo.Session, guildID string, memberID string) ([]string, error) {
    
    guild, _ := bot.State.Guild(guildID)
    
    member, err := bot.GuildMember(guildID, memberID)
    
    if err != nil {
        return nil, errors.New("Unknown Member")
    }
    
    perms := []string{}
    
    for _, role := range guild.Roles {
        if Includes(member.Roles, role.ID) == true {
            bitfield := structures.Bitfield{
                Bits: role.Permissions,
            }
            
            rperms := bitfield.Permissions()
            
            for _, flag := range rperms {
                if Includes(perms, flag) == false {
                    perms = append(perms, flag)
                }
            }
        }
    }
    
    return perms, nil
}