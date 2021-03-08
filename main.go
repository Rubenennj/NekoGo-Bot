package main

import (
    "./handlers"
    "github.com/bwmarrin/discordgo"
    "os"
    "./commands"
	"os/signal"
	"fmt"
	"syscall"
	"./config"
)

func main() {
    client, err := discordgo.New("Bot " + config.Token)
    if err != nil {
        fmt.Println("Failed to create bot session")
        
        return
    }
    
    //add msg event handler
    client.AddHandler(handlers.MessageCreate)
    
    err = client.Open()
    if err != nil {
        fmt.Println("Failed to connected to the gateway", err.Error())
        return
    }
    
    fmt.Println("Successfully logged in " + client.State.User.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
    
    fmt.Println(commands.Success)
    
    client.Close()
}