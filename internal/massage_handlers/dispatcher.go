package massagehandlers

import (
	"strings"

	"github.com/WhaleShip/BucketBot/api/router"
	api "github.com/WhaleShip/BucketBot/dto"
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
