package massagehandlers

import (
	"github.com/WhaleShip/BucketBot/api/router"
)

func handleStart(chatID int64) {
	router.SendMessage(chatID, "Привет!")
}
