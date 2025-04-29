package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Extract the message from the content without bot tag prefix and trailing white spaces
func ExtractedMessage(s *discordgo.Session, m *discordgo.MessageCreate) string {
	// Get the bot's user ID
	botUserID := s.State.User.ID

	botPrefix := "<@" + botUserID + ">"

	// Split the message to get the command
	message := m.Content[len(botPrefix):]

	// If the message is empty, then ignore
	if len(message) == 0 {
		return ""
	}

	// If the message has white space at the beginning, then remove it
	return strings.TrimSpace(message)
}
