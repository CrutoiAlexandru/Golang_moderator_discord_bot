package bot_control

import (
	"strings"

	"github.com/CrutoiAlexandru/Golang_moderator_discord_bot/config"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler in the main main function) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// test message
	if m.Content == "test" {
		for _, role := range m.Member.Roles {
			if role == "919917532167688213"{
				s.ChannelMessageSend(m.ChannelID, "halo")
			}
		}
	}

	if m.Content == "mod.channelcleanup" {
		
	}

	// delete the message and kick the user if it contains this set of words
	for _, word := range config.SECRET_WORDS {
		// check for the bad word
		if strings.Contains(m.Content, word) {
			// check the role of the user so we do not kick Admin users
			for _, role := range m.Member.Roles {
				if role == "919917532167688213" {
					// delete the message
					s.ChannelMessageDelete(m.ChannelID, m.ID)
					// send warning message
					s.ChannelMessageSend(m.ChannelID, "I know u are the Admin but that is not allowed!")
					return 
				}
			}
			// delete the message
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			// kick the user
			s.GuildMemberDelete(m.GuildID, m.Author.ID)
			// send message with reason
			s.ChannelMessageSend(m.ChannelID, "User " + m.Author.Username + " has been kicked for saying the wrong word!")
		}
	}
}