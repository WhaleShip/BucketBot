package bot_init

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/WhaleShip/BucketBot/config"
)

func SetWebhook(webhookURL string) error {
	cfg := config.GetConfig()
	resp, err := http.PostForm(fmt.Sprintf("%s%s/setWebhook", config.TelegramAPI, cfg.Bot.Token),
		url.Values{
			"url": {webhookURL},
		})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
