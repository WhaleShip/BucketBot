package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/WhaleShip/BucketBot/config"
	"github.com/WhaleShip/BucketBot/dto"
)

func SendMessage(chatID int64, text string, keyboard dto.InlineKeyboardMarkup) error {
	cfg := config.GetConfig()
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Bot.Token)
	message := dto.Response{Chat_id: chatID, Text: text, Markup: keyboard}
	body, err := json.Marshal(message)
	if err != nil {
		log.Fatal("error creating message: ", err)
	}
	_, err = http.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		log.Fatal("error sending message: ", err)
	}
	return err
}
