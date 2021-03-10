package main

import (
    "./events"
    "github.com/bwmarrin/discordgo"
    "os"
    "./commands"
	"os/signal"
	"./handlers"
	"fmt"
	"syscall"
	"./config"
)

var Uptime = 0 

func main() {
    client, err := discordgo.New("Bot " + config.Token)
    if err != nil {
        fmt.Println("Failed to create bot session")
        
        return
    }
    
    //add msg event handler
    client.AddHandler(events.MessageCreate)
    
    err = client.Open()
    if err != nil {
        fmt.Println("Failed to connected to the gateway", err.Error())
        return
    }
    
    commands.Load()
    
    fmt.Println("Successfully logged in " + client.State.User.Username)
    
    handlers.HandleStatus(client)
    
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
    
    client.Close()
}