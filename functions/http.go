package functions 

import "net/http"
import "time"
import "encoding/json"

var myClient = &http.Client{Timeout: 10 * time.Second}

func RequestJSON(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}