package bot_control

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// function made to remove set member
func remove_user(s *discordgo.Session, m *discordgo.MessageCreate, user string) {
	// get all the members in the guild
	guild_members, err := s.GuildMembers(m.Message.GuildID, "", 100)

	if err != nil {
		fmt.Println("Error retrieving guild members, ", err)
		return
	}

	// check our user against a list of members and remove it if it matches
	for _, member := range guild_members {
		if user == member.User.Username {
			// kick the user
			s.GuildMemberDelete(m.GuildID, member.User.ID)
			s.ChannelMessageSend(m.ChannelID, "Succesfully removed: " + member.User.Username)
			return
		}
	}

	// notiffy the Admin that the user was not found
	s.ChannelMessageSend(m.ChannelID, "Could not find set user, please see if the username is correct!")
}

func channel_cleanup(s *discordgo.Session, m *discordgo.MessageCreate) {
	// get the channel message structure for the last 100 message <<100 is the max number allowed by Discord>>
	channel_messages, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")
	channel_messages_ids := []string{}

	// retain the ids of the message types in a slice
	for _, messages := range channel_messages {
		channel_messages_ids = append(channel_messages_ids, messages.ID)
	}

	if err != nil {
		fmt.Println("Error retrieving channel messages ids, ", err)
		return 
	}

	// if the role of the message author is Admin continue deleting the latest 100 messages from current channel
	for _, role := range m.Member.Roles {
		if role == "919917532167688213"{
			s.ChannelMessagesBulkDelete(m.ChannelID, channel_messages_ids)
			return
		}
	}

	s.ChannelMessageSend(m.ChannelID, "You do not have Admin privileges.")
}

func secret_words(s *discordgo.Session, m *discordgo.MessageCreate, word string) {
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