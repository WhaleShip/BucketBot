package handler

import (
	"encoding/json"
	"log"
	"net/http"

	api "github.com/WhaleShip/BucketBot/dto"
	massagehandlers "github.com/WhaleShip/BucketBot/internal/massage_handlers"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var update api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Error while decoding update", err)
		return
	}
	go massagehandlers.HandleMessage(update)
}
