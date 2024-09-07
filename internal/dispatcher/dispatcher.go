package dispatcher

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database"
	"github.com/WhaleShip/BucketBot/internal/database/models"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/jackc/pgx/v5"
)

func directCallback(session *pgx.Conn, update dto.Update) error {
	var err error

	if update.CallbackQuery.Data == "create_note" {
		err = handleNewNoteCallback(update)
	} else if update.CallbackQuery.Data == "get_notes" {
		err = handleBackButton(update)
	}

	return err
}

func directMessage(session *pgx.Conn, update dto.Update) error {
	messageText := update.Message.Text
	userID := update.Message.Chat.ID
	var err error

	if strings.HasPrefix(messageText, "/start") {
		err = handleStart(update)
	} else if value, ok := state.GetUserState(update.Message.Chat.ID); ok && value == state.NewNoteState {
		err = handleNewNote(session, update)
	} else {
		n, err := database.GetNotesByUserID(session, userID)
		if err != nil {
			log.Println(err)
		}
		router.SendMessage(userID, NotesToString(n), nil)
	}

	return err
}
func NotesToString(notes []models.Note) string {
	var sb strings.Builder

	for _, note := range notes {
		// Создание строкового представления каждой заметки
		noteStr := fmt.Sprintf("Note ID: %d, Name: %s, Text: %s\n", note.ID, note.Name, note.Text)
		// Добавление этой строки в StringBuilder
		sb.WriteString(noteStr)
	}

	// Возвращение объединенной строки
	return sb.String()
}

func HandleMessage(session *pgx.Conn, update dto.Update) {
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
