package massagehandlers

import (
	"strings"

	api "github.com/WhaleShip/BucketBot/api/models"
	"github.com/WhaleShip/BucketBot/api/router"
)

func HandleMessage(update api.Update) {
	messageText := update.Message.Text
	chatID := update.Message.Chat.ID

	if strings.HasPrefix(messageText, "/start") {
		handleStart(chatID)
	} else {
		router.SendMessage(chatID, messageText)
	}
}
