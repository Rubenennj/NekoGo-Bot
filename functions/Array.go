package functions

import "strings"

func Goof(array []string) []string {
    for i, value := range array {
        array[i] = strings.Title(strings.Replace(strings.ToLower(value), "_", " ", -1))
    }
    return array
}

func Includes(array []string, value string) bool {
    for _, val := range array {
        if val == value {
            return true
        }
    }
    
    return false
}