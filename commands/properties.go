package commands 

import "./developer"
import "fmt"
import "./info"
import "./staff"
import "./fun"

func Load() {
    fmt.Println(developer.Success)
    fmt.Println(info.Success)
    fmt.Println(staff.Success)
    fmt.Println(fun.Success)
}