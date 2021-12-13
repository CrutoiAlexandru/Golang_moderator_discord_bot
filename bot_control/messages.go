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

	// command for Admin to remove certain member to guild
	if strings.HasPrefix(m.Content, "mod.rm") {
		// isolate the member we want to remove
		user := strings.Replace(m.Content, "mod.rm", "", -1)
		user = strings.Replace(user, " ", "", -1)

		// check if command member is Admin
		for _, role := range m.Member.Roles {
			if role == "919917532167688213"{
				remove_user(s, m, user)
				return
			}
		}

		s.ChannelMessageSend(m.ChannelID, "You do not have Admin privileges.")
	}

	// test message
	if m.Content == "mod.add" {
		for _, role := range m.Member.Roles {
			if role == "919917532167688213"{
				s.ChannelMessageSend(m.ChannelID, "halo")
			}
		}
	}

	// Admin command to delete max 100 messages in a channel
	if m.Content == "mod.channelcleanup" {
		// it is usefull to be a goroutine as sometimes it takes a bit long and we do not want users to wait for it 
		go channel_cleanup(s, m)
	}

	// delete the message and kick the user if it contains this set of words
	for _, word := range config.SECRET_WORDS {
		secret_words(s, m, word)
	}
}