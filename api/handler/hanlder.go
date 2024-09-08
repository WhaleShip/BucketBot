package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WhaleShip/BucketBot/config"
	api "github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/app"
	"github.com/jackc/pgx/v5"
)

func WebhookHandler(session *pgx.Conn, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered from panic:", err)
		}
	}()

	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != config.GetConfig().Webhook.Secret {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var update api.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Error while decoding update", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))

	app.HandleUpdate(session, &update)
}
