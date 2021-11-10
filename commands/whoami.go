package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func WhoAmi(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Get the "member" of the user.
	// This contains information about the user in the specific guild
	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return err
	}

	// Get all roles that are available on the guild
	guildRoles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		return err
	}

	// Compare all guild roles with the roles the user has
	roleString := ""
	for _, role := range member.Roles {
		for _, guildRole := range guildRoles {
			if role == guildRole.ID {
				roleString += guildRole.Mention() + " "
			}
		}
	}
	// If the user is in no guild, set this replacement
	if roleString == "" {
		roleString = "_None_"
	} else {
		// Strip the last space
		roleString = roleString[:len(roleString)-1]
	}

	joinDate, err := member.JoinedAt.Parse()
	if err != nil {
		return err
	}

	embed := discordgo.MessageEmbed{
		Title: fmt.Sprintf("User: %s#%s", m.Author.Username, m.Author.Discriminator),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Username",
				Value:  m.Author.Username,
				Inline: true,
			},
			{
				Name:   "Discriminator",
				Value:  m.Author.Discriminator,
				Inline: true,
			},
			{
				Name: "Join Date",
				// This is Discord specific formatting
				// See https://discord.com/developers/docs/reference#message-formatting
				Value:  fmt.Sprintf("<t:%d:F>", joinDate.Unix()),
				Inline: true,
			},
			{
				Name:   "Roles",
				Value:  roleString,
				Inline: false,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL("512")},
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embed)

	return err
}
