package dispatcher

import (
	"log"
	"strings"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/state"
)

func directCallback(update dto.Update) {
	if update.CallbackQuery.Data == "create_note" {
		handleNewNoteCallback(update)
	} else {
		log.Print("Error: nknown Callback recieved")
	}
}

func directMessage(update dto.Update) {
	messageText := update.Message.Text
	chatID := update.Message.Chat.ID

	if strings.HasPrefix(messageText, "/start") {
		handleStart(update)
	} else if value, ok := state.GetState(update.Message.Chat.ID); ok && value == state.NewNoteState {
		log.Printf("done")
	} else {
		router.SendMessage(chatID, messageText, nil)
	}
}

func HandleMessage(update dto.Update) {

	if update.CallbackQuery != nil {
		directCallback(update)

	} else if update.Message != nil {
		directMessage(update)
	}
}
