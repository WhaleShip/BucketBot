package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const TelegramAPI = "https://api.telegram.org/bot"

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

var (
	config *Config
	once   sync.Once
)

func LoadJsonConfig(file string) (*Config, error) {
	var err error
	once.Do(func() {
		file, err := os.Open(file)
		if err != nil {
			log.Fatalf("Failed to open config file: %v", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		config = &Config{}
		if err := decoder.Decode(config); err != nil {
			log.Fatalf("Failed to decode config file: %v", err)
		}
	})
	return config, err
}

func GetConfig() *Config {
	return config
}
