package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/WhaleShip/BucketBot/api/handler"
	config "github.com/WhaleShip/BucketBot/config/app"
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

	cfg, err := config.LoadJsonConfig("config/app/app_config.json")
	if err != nil {
		log.Fatal("Error loading config: ", err)
		return
	}

	state.InitializeStateMachine()

	conn, err := database.GetInitializedDb()
	if err != nil {
		log.Fatal("Error connection DB: ", err)
	}
	defer conn.Close(context.Background())

	http.HandleFunc(cfg.Webhook.Path, handler.WebhookHandler)

	if err := bot_init.SetWebhook(cfg.Webhook.Host + cfg.Webhook.Path); err != nil {
		log.Fatal("Error setting webhook: ", err)
		return
	}

	log.Printf("Starting server on port %d", cfg.Webapp.Port)
	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Webapp.Port), nil); err != nil {
			log.Fatal("Error on start up: ", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
