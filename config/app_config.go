package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const TelegramAPI = "https://api.telegram.org/bot"

type Config struct {
	Webhook WebhookConf `json:"webhook"`
	Bot     BotConf     `json:"bot"`
	Webapp  WebappConf  `json:"webapp"`
}

type WebhookConf struct {
	Host   string `json:"host"`
	Path   string `json:"path"`
	Secret string `json:"secret"`
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
)

func LoadJsonConfig(file string) (*Config, error) {
	jsonCfg, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer jsonCfg.Close()

	decoder := json.NewDecoder(jsonCfg)
	config = &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}
	return config, err
}

func GetConfig() *Config {
	return config
}
