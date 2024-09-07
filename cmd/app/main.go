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
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func HandlerWithDbConnection(session *pgx.Conn, handler func(*pgx.Conn, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(session, w, r)
	}
}

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

	session, err := database.GetInitializedDb()
	if err != nil {
		log.Fatal("Error connection DB: ", err)
	}
	defer session.Close(context.Background())

	http.HandleFunc(cfg.Webhook.Path, HandlerWithDbConnection(session, handler.WebhookHandler))

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
