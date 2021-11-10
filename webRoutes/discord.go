package webRoutes

import (
	"fmt"
	"net/http"
	"strings"
	"webserver/utils"
)

func Discord(w http.ResponseWriter, r *http.Request) {

	var messages []string

	{
		utils.GetBridge().Mutex.RLock()
		defer utils.GetBridge().Mutex.RUnlock()

		for _, msg := range utils.GetBridge().Messages {
			messages = append(messages, fmt.Sprintf("Guild: %s, User: %s, Username: %s#%s, Message: %s", msg.GuildID, msg.UserID, msg.Username, msg.Discriminator, msg.Message))
		}
	}
	_, _ = fmt.Fprintf(w, "Discord Message Log: \n%s", strings.Join(messages, "\n"))
}
