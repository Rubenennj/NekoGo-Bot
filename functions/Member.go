package functions

import (
    "github.com/bwmarrin/discordgo"
    "../utils"
)

func MemberHighestRole (bot *discordgo.Session, guildID string, member *discordgo.Member) (*discordgo.Role, error) {
    //initialize a role
    role := &discordgo.Role{}
    //loop over member roles
    for _, id := range member.Roles {
        //request role
        r, err := GetRole(bot, guildID, id)
        
        if err != nil {
            return nil, err 
        }
        
        //check if the position is higher
        if r.Position > role.Position && r.Color != 0 {
            role = r 
        }
    }
    
    return role, nil 
}

func MemberIsStaff (member *discordgo.Member) bool {
    for _, roleID := range member.Roles {
        for _, id := range utils.StaffRoles {
            if id == roleID {
                return true 
            }
        }
    }
    
    return false
}