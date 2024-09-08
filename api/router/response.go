package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	config "github.com/WhaleShip/BucketBot/config/app"
	"github.com/WhaleShip/BucketBot/dto"
)

func sendUpdate(url string, updateBody any) error {

	body, err := json.Marshal(updateBody)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected response from Telegram API: %s", string(bodyBytes))
	}

	return err
}

func SendMessage(chatID int, text string, keyboard *dto.InlineKeyboardMarkup) error {
	sendMessageURL := "https://api.telegram.org/bot%s/sendMessage"
	cfg := config.GetConfig()
	url := fmt.Sprintf(sendMessageURL, cfg.Bot.Token)

	message := &dto.ResponseMessage{Chat_id: chatID, Text: text, Markup: keyboard}
	err := sendUpdate(url, message)
	if err != nil {
		log.Print("error sending message: ", err)
	}

	return err
}

func CallbackEditMessage(chatID int, messageID int, newText string, newMarkup *dto.InlineKeyboardMarkup) error {
	editMEssageUrl := "https://api.telegram.org/bot%s/editMessageText"

	cfg := config.GetConfig()
	url := fmt.Sprintf(editMEssageUrl, cfg.Bot.Token)

	newMessage := &dto.ResponseMessage{Chat_id: chatID, Text: newText, Markup: newMarkup}
	update := &dto.ResponseEditMessage{ResponseMessage: newMessage, Message_id: int64(messageID)}
	err := sendUpdate(url, update)
	return err
}

func DeleteWebhook() error {
	deleteWebhookURL := "https://api.telegram.org/bot%s/deleteWebhook"

	cfg := config.GetConfig()
	url := fmt.Sprintf(deleteWebhookURL, cfg.Bot.Token)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete webhook, status code: %d", resp.StatusCode)
	}

	return nil
}
