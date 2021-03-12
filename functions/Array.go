package functions

import "strings"
import "sort"

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
    arr := []string{}
    for _, val := range array {
        arr = append(arr, fn(val))
    }
  
    return arr 
}

func Join(array []string, sep string) string {
    return strings.Join(array, sep)
}

func Plural (argument string, n int64) string {
    plr := ""
    
    if n > 1 {
        plr = "s"
    }
    
    return argument + plr 
}

func SortMap (mapper map[string]int) map[string]int {
    type temp struct {
        Key string
        Value int 
    }
    
    var ss []temp 
    
    current := map[string]int{}
    
    for k, v := range mapper {
        ss = append(ss, temp{k, v})
    }
    
    sort.Slice(ss, func(x, y int) bool {
        return ss[x].Value > ss[y].Value 
    })
    
    
    dat := make([]string, len(mapper))
    
    for i, kv := range ss {
        dat[i] = kv.Key
    }
    
    for _, k := range dat {
        current[k] = mapper[k]
    }
    
    return current
}

func SortMapStringInt(values map[string]int) []string {
    type kv struct {
        Key   string
        Value int
    }
    var ss []kv
    for k, v := range values {
        ss = append(ss, kv{k, v})
    }
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })
    ranked := make([]string, len(values))
    for i, kv := range ss {
        ranked[i] = kv.Key
    }
    return ranked
}