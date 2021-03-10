package functions

import "strings"

func Has(array []string, flag string) bool {
    for _, val := range array {
        if val == flag {
            return true
        }
    }
    
    return false
}

func ArrayMapIncludes(array []string, mapper map[string]string) bool {
    for _, value := range array {
        for _, id := range mapper {
            if value == id {
                return true
            }
        }
    }
    
    return false
}

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

func Map(array []string, fn func(str string) string) []string{
    for i, val := range array {
        array[i] = fn(val)
    }
  
    return array
}

func Join(array []string, sep string) string {
    return strings.Join(array, sep)
}