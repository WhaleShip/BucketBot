package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/WhaleShip/BucketBot/api/handler"
	"github.com/WhaleShip/BucketBot/config"
	"github.com/WhaleShip/BucketBot/internal/database"
	bot_init "github.com/WhaleShip/BucketBot/internal/init"
	"github.com/WhaleShip/BucketBot/internal/state"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg, err := config.LoadJsonConfig("config/config.json")
	if err != nil {
		log.Fatal("Error loading config: ", err)
		return
	}

	state.InitializeStateMachine()

	_, _ = database.GetInitializedDb()

	http.HandleFunc(cfg.Webhook.Path, handler.WebhookHandler)

	if err := bot_init.SetWebhook(cfg.Webhook.Host + cfg.Webhook.Path); err != nil {
		log.Fatal("Error setting webhook: ", err)
		return
	}

	log.Printf("Starting server on port %d", cfg.Webapp.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Webapp.Port), nil); err != nil {
		log.Fatal("Error on start up: ", err)
	}
}
