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
	"github.com/WhaleShip/BucketBot/api/router"
	"github.com/WhaleShip/BucketBot/config"
	"github.com/WhaleShip/BucketBot/internal/database"
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
		log.Fatal("error loading .env file: ", err)
	}

	cfg, err := config.LoadJsonConfig("config/app_config.json")
	if err != nil {
		log.Fatal("error loading config: ", err)
		return
	}

	state.InitializeStateMachine()

	session, err := database.GetInitializedDb()
	if err != nil {
		log.Fatal("error connection DB: ", err)
	}
	defer session.Close(context.Background())

	http.HandleFunc(cfg.Webhook.Path, HandlerWithDbConnection(session, handler.WebhookHandler))

	if err := router.SetWebhook(cfg.Webhook.Host + cfg.Webhook.Path); err != nil {
		log.Fatal("error setting webhook: ", err)
		return
	}
	defer func() {
		err = router.DeleteWebhook()
		if err != nil {
			log.Println("error deleting webhook: ", err)
		} else {
			log.Println("Webhook deleted successfuly!")
		}

	}()

	log.Printf("Starting server on port %d", cfg.Webapp.Port)
	go func() {
		if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Webapp.Port), nil); err != nil {
			log.Fatal("error on start up: ", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
