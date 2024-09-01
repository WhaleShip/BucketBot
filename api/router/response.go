package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/WhaleShip/BucketBot/config"
)

func SendMessage(chatID int64, text string) error {
	cfg := config.GetConfig()
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Bot.Token)
	message := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	body, _ := json.Marshal(message)
	_, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	return err
}
