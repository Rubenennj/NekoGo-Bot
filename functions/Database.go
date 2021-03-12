package functions

import "../config"
import "encoding/json"
import "github.com/bwmarrin/discordgo"

type Data struct {
    MessageCount int `json:"message_count,omitempty"`
    Messages map[string]int `json:"messages"`
}

func (data *Data) Save(key string) error {
    err := DatabaseSet(key, data)
    if err != nil {
        return err
    }
    
    return nil
}

func DatabaseGet (key string) (*Data, error) {
    bytes, err := config.Database.Get([]byte(key))
    
    if err != nil {
        return nil, err 
    }
    
    data := Data{} 
    
    err = json.Unmarshal(bytes, &data)
    
    if err != nil {
        return nil, err
    }
    
    return &data, nil
}

func DatabaseMustGet (key string) (*Data, error) {
    data, err := DatabaseGet(key)
    
    //default data
    defaultData := Data{
        MessageCount: 0,
        Messages: map[string]int{},
    }
    
    //key does not exist
    if err != nil {
        err = DatabaseSet(key, &defaultData)
        if err != nil {
            return nil, err
        } else {
            return &defaultData, nil
        }
    }
    
    return data, nil
}

func DatabaseSet (key string, data *Data) error {
    compact, err := json.Marshal(&data)
    if err != nil {
        return err
    }
    
    err = config.Database.Put([]byte(key), compact)
    if err != nil {
        return err
    }
    
    return nil
}

func AddMessagesWithData (key string, value int, msg *discordgo.MessageCreate) error {
    data, err := DatabaseMustGet(key)
    if err != nil {
        return err
    }
    
    data.MessageCount += value 
    
    _, exists := data.Messages[msg.ChannelID]
    
    if exists == false {
        data.Messages[msg.ChannelID] = 0 
    }
    
    data.Messages[msg.ChannelID] = data.Messages[msg.ChannelID] + value
    
    err = DatabaseSet(key, data)
    if err != nil {
        return err 
    }
    
    return nil 
}

func AddMessages (key string, value int) error {
    data, err := DatabaseMustGet(key)
    if err != nil {
        return err
    }
    
    data.MessageCount += value 
    
    err = DatabaseSet(key, data)
    if err != nil {
        return err 
    }
    
    return nil 
}

func DatabaseSortByCount (mapper map[string]*Data) []string {
    intMap := map[string]int{}
    
    for str, data := range mapper {
        intMap[str] = data.MessageCount
    }
    
    array := SortMapStringInt(intMap)
    
    return array 
}

func DatabaseGetKeysWith (k string) (map[string]*Data, error) {
    byted := []byte(k)
    
    m := map[string]*Data{}
    
    err := config.Database.Scan(byted, func (key []byte) error {
        data, err := DatabaseGet(string(key))
        if err != nil {
            return err 
        }
        
        m[string(key)] = data 
        
        return nil 
    })
    
    if err != nil {
        return nil, err 
    }
    
    return m, nil 
}