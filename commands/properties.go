package commands 

import "./developer"
import "fmt"
import "./info"

func Load() {
    fmt.Println(developer.Success)
    fmt.Println(info.Success)
}