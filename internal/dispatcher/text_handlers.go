package dispatcher

import (
	"fmt"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/jackc/pgx/v5"
)

func handleNewNote(session *pgx.Conn, update dto.Update) error {
	noteText := update.Message.Text
	if noteText == "" {
		router.SendMessage(update.Message.Chat.ID, "Полученное сообщение некоректно", nil)
	}

	err := database.AddNewNote(session, noteText, update.Message.Chat.ID)
	if err != nil {
		return fmt.Errorf("error while adding note to db: %w", err)

	}

	err = router.SendMessage(update.Message.Chat.ID, "заметка сохранена!", markups.GoBackKeyboard)
	if err != nil {
		return fmt.Errorf("error sending OK message: %w", err)
	}

	state.SetUserState(update.Message.Chat.ID, state.NoState)

	return nil
}
