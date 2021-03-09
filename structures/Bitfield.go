package structures

import "../utils"

type Bitfield struct {
    Bits int64
}

func (b Bitfield) Has(perm string) bool {
    perms := b.Permissions()
    
    for _, flag := range perms {
        if flag == perm {
            return true
        }
    }
    
    return false
}

func (b Bitfield) Permissions() []string {
    perms := []string{}
    
    for key, value := range utils.Permissions {
        if (b.Bits & value) == value {
            if key == "ADMINISTRATOR" {
                perms = []string{}
                for key, _ := range utils.Permissions {
                    perms = append(perms, key)
                }
                return perms
            }
            perms = append(perms, key)
        }
    }
    
    return perms
}