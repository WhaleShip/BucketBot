package app

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

func handleNewNoteCallback(update *dto.Update) error {
	if update.CallbackQuery.From == nil {
		return errors.New("error with callback format: no information about message")
	} else {
		err := router.EditMessage(update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID, "Отправь сообщение, и оно станет новой заметкой!", markups.GoBackKeyboard)
		state.SetUserState(update.CallbackQuery.From.ID, state.NewNoteState)

		if err != nil {
			return fmt.Errorf("error sending callback answer: %w", err)
		}
	}

	return nil
}

func safecheckCallbackWithParam(update *dto.Update) (int, error) {
	if update.CallbackQuery.From == nil {
		return 0, errors.New("error with callback format: no information about message")
	}

	fields := strings.Fields(update.CallbackQuery.Data)
	if len(fields) != 2 {
		return 0, fmt.Errorf("error reading callback data: %s", update.CallbackQuery.Data)
	}
	param, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, fmt.Errorf("error reading offset in callback data: %s", update.CallbackQuery.Data)
	}

	return param, nil
}

func handleGetNoteListCallback(session *pgx.Conn, update *dto.Update) error {
	offset, err := safecheckCallbackWithParam(update)
	if err != nil {
		return err
	}

	user := update.CallbackQuery.From
	notes, err := database.GetSomeUserNotes(session, user.ID, markups.NotesPerScreenCount, offset)
	if err != nil {
		return err
	}

	err = router.EditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, texts.MainText, markups.GetNotesKeyboard(notes, offset))
	if err != nil {
		return fmt.Errorf("error sending callback answer: %w", err)
	}
	state.SetUserState(update.CallbackQuery.From.ID, state.NoState)

	return nil
}

func handleGetNoteCallback(session *pgx.Conn, update *dto.Update) error {
	noteID, err := safecheckCallbackWithParam(update)
	if err != nil {
		return err
	}

	user := update.CallbackQuery.From
	note, err := database.GetNoteByIDForOwner(session, noteID, user.ID)
	if err != nil {
		return fmt.Errorf("error getting note: %s", update.CallbackQuery.Data)
	}

	err = router.EditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, note.Text, markups.GetNoteGoBackKeyboard(noteID))
	if err != nil {
		return fmt.Errorf("error sending callback answer: %w", err)
	}

	return nil
}

func deleteNoteHandler(session *pgx.Conn, update *dto.Update) error {
	noteID, err := safecheckCallbackWithParam(update)
	if err != nil {
		return err
	}

	user := update.CallbackQuery.From
	err = database.DeleteNoteByIDByOwner(session, noteID, user.ID)
	if err != nil {
		return fmt.Errorf("error getting note: %s", update.CallbackQuery.Data)
	}

	err = router.EditMessage(update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID, texts.NoteDeletedText, markups.GoBackKeyboard)
	if err != nil {
		return fmt.Errorf("error sending callback answer: %w", err)
	}

	return nil
}
