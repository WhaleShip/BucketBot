package dispatcher

import (
	"log"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/assets/texts"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/state"
)

func handleNewNoteCallback(update dto.Update) {
	if update.CallbackQuery.From == nil {
		log.Print("Error with callback format: no information about message")
	} else {
		err := router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID, "Отправь сообщение, и оно станет новой заметкой!", markups.GoBackKeyboard)
		state.SetUserState(update.CallbackQuery.From.ID, state.NewNoteState)

		if err != nil {
			log.Print("Error sending callback answer: ", err)
		}
	}
}

func handleBackButton(update dto.Update) {
	if update.CallbackQuery.From == nil {
		log.Print("Error with callback format: no information about message")
	} else {
		err := router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID, texts.MainText, markups.GetNotesKeyboard())
		state.SetUserState(update.CallbackQuery.From.ID, state.NoState)

		if err != nil {
			log.Print("Error sending callback answer: ", err)
		}
	}
}
