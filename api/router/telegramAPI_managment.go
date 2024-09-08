package router

import (
	"fmt"
	"net/http"

	"github.com/WhaleShip/BucketBot/config"
	"github.com/WhaleShip/BucketBot/dto"
)

func SetWebhook(webhookURL string) error {
	cfg := config.GetConfig()

	url := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", cfg.Bot.Token)
	payload := dto.WebhookRequest{
		URL:         webhookURL,
		SecretToken: cfg.Webhook.Secret,
	}

	err := sendUpdate(url, payload)

	if err != nil {
		return err
	}

	return nil
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
