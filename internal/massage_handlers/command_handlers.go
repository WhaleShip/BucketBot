package massagehandlers

import (
	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/internal/markups"
)

func handleStart(chatID int64) {
	router.SendMessage(chatID, "Привет!", markups.GetNotesKeyboard())
}
