package dispatcher

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/assets/texts"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/jackc/pgx/v5"
)

const notesPerScreenCount int = 6

func handleNewNoteCallback(update *dto.Update) error {
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

func handleGetNoteListCallback(session *pgx.Conn, update *dto.Update) error {
	user := update.CallbackQuery.From
	if user == nil {
		return errors.New("error with callback format: no information about message")
	}

	fields := strings.Fields(update.CallbackQuery.Data)
	if len(fields) != 2 {
		return fmt.Errorf("error reading callback data: %s", update.CallbackQuery.Data)
	}
	offset, err := strconv.Atoi(fields[1])
	if err != nil {
		return fmt.Errorf("error reading offset in callback data: %s", update.CallbackQuery.Data)
	}

	notes, err := database.GetSomeUserNotes(session, user.ID, notesPerScreenCount, offset)
	if err != nil {
		return err
	}

	err = router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, texts.MainText, markups.GetNotesKeyboard(notes))
	state.SetUserState(update.CallbackQuery.From.ID, state.NoState)

	if err != nil {
		return fmt.Errorf("error sending callback answer: %w", err)
	}

	return nil
}

func handleGetNoteCallback(session *pgx.Conn, update *dto.Update) error {
	user := update.CallbackQuery.From
	if user == nil {
		return errors.New("error with callback format: no information about message")
	}
	fields := strings.Fields(update.CallbackQuery.Data)
	if len(fields) != 2 {
		return fmt.Errorf("error reading callback data: %s", update.CallbackQuery.Data)
	}
	noteID, err := strconv.Atoi(fields[1])
	if err != nil {
		return fmt.Errorf("error reading noteID in callback data: %s", update.CallbackQuery.Data)
	}

	note, err := database.GetNoteByID(session, noteID)
	if err != nil {
		return fmt.Errorf("error getting note: %s", update.CallbackQuery.Data)

	}

	err = router.CallbackEditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, note.Text, markups.GoBackKeyboard)

	if err != nil {
		return fmt.Errorf("error sending callback answer: %w", err)
	}

	return nil
}
