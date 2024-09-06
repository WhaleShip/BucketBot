package dispatcher

import (
	"log"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/dto"
)

func handleNewNoteCallback(update dto.Update) {
	if update.CallbackQuery.Message == nil {
		log.Print("Error with callback format: no information about message")
	}
	err := router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, "Отправь сообщение, и оно станет новой заметкой!")

	if err != nil {
		log.Print("Error sending callback answer: ", err)
	}
}
