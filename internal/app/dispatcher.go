package app

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/jackc/pgx/v5"
)

func directCallback(session *pgx.Conn, update *dto.Update) error {
	var err error
	callbackQueryText := update.CallbackQuery.Data

	if callbackQueryText == markups.CreateNoteCallback {
		err = handleNewNoteCallback(update)
	} else if strings.HasPrefix(callbackQueryText, markups.GetNoteListCallback) {
		err = handleGetNoteListCallback(session, update)
	} else if strings.HasPrefix(callbackQueryText, markups.GetNoteCallback) {
		err = handleGetNoteCallback(session, update)
	} else if callbackQueryText == markups.NoPageCallback {
	} else {
		err = fmt.Errorf("unknown callback recieved: %s", callbackQueryText)
	}

	return err
}

func directMessage(session *pgx.Conn, update *dto.Update) error {
	var err error
	messageText := update.Message.Text

	if strings.HasPrefix(messageText, "/start") {
		err = handleStart(session, update)
	} else if value, ok := state.GetUserState(update.Message.Chat.ID); ok && value == state.NewNoteState {
		err = handleNewNote(session, update)
	} else {
		handleUselessText(update)
	}

	return err
}

func HandleUpdate(session *pgx.Conn, update *dto.Update) {
	var err error

	if update.CallbackQuery != nil {
		err = directCallback(session, update)

	} else if update.Message != nil {
		err = directMessage(session, update)
	} else {
		err = errors.New("ivnalid update")
	}
	if err != nil {
		log.Print("error parsing update: ", err)
	}

}
