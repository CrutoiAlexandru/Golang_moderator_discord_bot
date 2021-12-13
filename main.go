package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/CrutoiAlexandru/Golang_moderator_discord_bot/config"
	"github.com/CrutoiAlexandru/Golang_moderator_discord_bot/bot_control"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// create a Discord session using the provided token
	discord, err := discordgo.New("Bot " +  config.TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session, ", err)
		return	
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(bot_control.MessageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}