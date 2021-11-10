package utils

import (
	"sync"
)

var bridgeSingleton *Bridge

type Bridge struct {
	Messages []DiscordMessage
	Mutex    sync.RWMutex
}

type DiscordMessage struct {
	GuildID       string
	UserID        string
	Username      string
	Discriminator string
	Message       string
}

func GetBridge() *Bridge {
	if bridgeSingleton == nil {
		bridgeSingleton = &Bridge{
			Messages: []DiscordMessage{},
			Mutex:    sync.RWMutex{},
		}
	}

	return bridgeSingleton
}

func (b *Bridge) AddMessage(msg DiscordMessage) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Messages = append(b.Messages, msg)
}
