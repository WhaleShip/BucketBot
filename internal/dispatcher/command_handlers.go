package dispatcher

import (
	"errors"
	"fmt"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/assets/texts"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/jackc/pgx/v5"
)

func handleStart(session *pgx.Conn, update *dto.Update) error {
	chat := update.Message.Chat
	if chat == nil {
		return errors.New("error with callback format: no information about message")
	}

	notes, err := database.GetSomeUserNotes(session, chat.ID, markups.NotesPerScreenCount, 0)
	if err != nil {
		return fmt.Errorf("error GettingNotes: %w", err)
	}

	markup := markups.GetNotesKeyboard(notes, 0)
	err = router.SendMessage(update.Message.Chat.ID, "Привет!\n"+texts.MainText, markup)

	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}

	state.SetUserState(update.Message.Chat.ID, state.NoState)

	return nil
}
