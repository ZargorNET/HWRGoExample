package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
	"webserver/commands"
	"webserver/utils"
)

func onReady(s *discordgo.Session, event *discordgo.Event) {
	status := []string{"KursB > KursA", "KursA > KursB", "Ich bin eine Biene"}
	index := 0

	go func() {
		for {
			_ = s.UpdateGameStatus(0, status[index])
			index = (index + 1) % len(status)
			<-time.NewTimer(10 * time.Second).C
		}
	}()
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore Bot messages
	if m.Author.Bot {
		return
	}
	log.Printf("Received message from %s#%s: %s", m.Author.Username, m.Author.Discriminator, m.Content)
	utils.GetBridge().AddMessage(utils.DiscordMessage{
		GuildID:       m.GuildID,
		UserID:        m.Author.ID,
		Username:      m.Author.Username,
		Discriminator: m.Author.Discriminator,
		Message:       m.Content,
	})

	command := m.Content

	// Check if the message has our prefix, else ignore.
	if !strings.HasPrefix(command, "+") {
		return
	}

	// Remove the first character (the prefix) and lower case the entire text
	command = strings.ToLower(command[1:])

	var err error

	switch command {
	case "hello":
		err = commands.HelloCommand(s, m)
		break
	case "whoami":
		err = commands.WhoAmi(s, m)
		break
	}

	if err == nil {
		return
	}

	log.Printf("Error while executing command: %s: %v", command, err)

	// Send error embed
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Something went wrong",
		Description: "We apologize.",
		Color:       0xFF0000,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: "https://ih1.redbubble.net/image.370389938.3139/flat,750x,075,f-pad,750x1000,f8f8f8.jpg"},
	})
}
