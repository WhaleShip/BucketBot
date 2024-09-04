package massagehandlers

import (
	"strings"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/dto"
)

func HandleMessage(update dto.Update) {
	messageText := update.Message.Text
	chatID := update.Message.Chat.ID

	if strings.HasPrefix(messageText, "/start") {
		handleStart(chatID)
	} else {
		router.SendMessage(chatID, messageText, dto.InlineKeyboardMarkup{})
	}
}
