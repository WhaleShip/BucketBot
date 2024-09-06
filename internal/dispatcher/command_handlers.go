package dispatcher

import (
	"log"

	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/assets/markups"
	"github.com/WhaleShip/BucketBot/assets/texts"
	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/state"
)

func handleStart(update dto.Update) {
	markup := markups.GetNotesKeyboard()
	err := router.SendMessage(update.Message.Chat.ID, "Привет!\n"+texts.MainText, markup)

	if err != nil {
		log.Print("Error sending message: ", err)
	} else {
		state.SetUserState(update.Message.Chat.ID, state.NoState)
	}
}
