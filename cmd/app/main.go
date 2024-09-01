package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WhaleShip/BucketBot/api/handler"
	"github.com/WhaleShip/BucketBot/config"
	"github.com/WhaleShip/BucketBot/internal/database"
	bot_init "github.com/WhaleShip/BucketBot/internal/init"
)

func main() {
	cfg, err := config.LoadJsonConfig("config/config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	_, err = database.ConnectPostgres(database.Config{
		Host:     "db",
		Port:     "5432",
		Username: "user",
		Password: "pass",
		DBName:   "buckets",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("db connection fail: %s", err.Error())
	}

	http.HandleFunc(cfg.Webhook.Path, handler.WebhookHandler)

	if err := bot_init.SetWebhook(cfg.Webhook.Host + cfg.Webhook.Path); err != nil {
		fmt.Println("Error setting webhook:", err)
		return
	}

	log.Printf("Starting server on port %d", cfg.Webapp.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.Webapp.Port), nil); err != nil {
		log.Fatal("Error on start up:", err)
	}
}
