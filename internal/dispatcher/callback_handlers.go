package dispatcher

import (
	"errors"
	"fmt"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/assets/texts"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/state"
)

func handleNewNoteCallback(update dto.Update) error {
	if update.CallbackQuery.From == nil {
		return errors.New("error with callback format: no information about message")
	} else {
		err := router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID, "Отправь сообщение, и оно станет новой заметкой!", markups.GoBackKeyboard)
		state.SetUserState(update.CallbackQuery.From.ID, state.NewNoteState)

		if err != nil {
			return fmt.Errorf("error sending callback answer: %w", err)
		}
	}

	return nil
}

func handleBackButton(update dto.Update) error {
	if update.CallbackQuery.From == nil {
		return errors.New("error with callback format: no information about message")
	} else {
		err := router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID, texts.MainText, markups.GetNotesKeyboard())
		state.SetUserState(update.CallbackQuery.From.ID, state.NoState)

		if err != nil {
			return fmt.Errorf("error sending callback answer: %w", err)
		}
	}

	return nil
}
