package handler

import (
	"encoding/json"
	"log"
	"net/http"

	api "github.com/WhaleShip/BucketBot/dto"
	massagehandlers "github.com/WhaleShip/BucketBot/internal/dispatcher"
	"github.com/jackc/pgx/v5"
)

func WebhookHandler(session *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	var update api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Error while decoding update", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))

	massagehandlers.HandleMessage(session, update)
}
