package bot_control

import (
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

	// Admin command to delete max 100 messages in a channel
	if m.Content == "mod.channelcleanup" {
		channel_cleanup(s, m)
	}

	// delete the message and kick the user if it contains this set of words
	for _, word := range config.SECRET_WORDS {
		secret_words(s, m, word)
	}
}