package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var botToken string = ""

const telegramAPI = "https://api.telegram.org/bot"

type Config struct {
	Webhook WebhookConf `json:"webhook"`
	Bot     BotConf     `json:"bot"`
	Webapp  WebappConf  `json:"webapp"`
}

type WebhookConf struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type BotConf struct {
	Token string `json:"token"`
}

type WebappConf struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

type Chat struct {
	ID int64 `json:"id"`
}

func sendMessage(chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	message := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	body, _ := json.Marshal(message)
	_, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	return err
}

func handleStart(chatID int64) {
	sendMessage(chatID, "Привет!")
}

func handleMessage(update Update) {
	messageText := update.Message.Text
	chatID := update.Message.Chat.ID

	if strings.HasPrefix(messageText, "/start") {
		handleStart(chatID)
	} else {
		sendMessage(chatID, messageText)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Error while decoding update", err)
		return
	}
	handleMessage(update)
}

func setWebhook(botToken, webhookURL string) error {
	resp, err := http.PostForm(fmt.Sprintf("%s%s/setWebhook", telegramAPI, botToken),
		url.Values{
			"url": {webhookURL},
		})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func loadJsonConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := loadJsonConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	botToken = config.Bot.Token

	http.HandleFunc(config.Webhook.Path, webhookHandler)

	if err := setWebhook(botToken, config.Webhook.Host+config.Webhook.Path); err != nil {
		fmt.Println("Error setting webhook:", err)
		return
	}

	log.Printf("Starting server on port %d", config.Webapp.Port)
	if err := http.ListenAndServe(":"+strconv.Itoa(config.Webapp.Port), nil); err != nil {
		log.Fatal("Error on start up:", err)
	}
}
