package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HelloCommand(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Hello World!")

	return err
}
