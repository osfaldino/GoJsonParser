package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName string   `json:"appName"`
	Version string   `json:"version"`
	Debug   bool     `json:"debug"`
	Users   []string `json:"users"`
}

func (cfg *Config) ParseConfig(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var cfg Config
	err := cfg.ParseConfig("cfg.json")
	if err != nil {
		fmt.Printf("Ошибка: %s", err)
		os.Exit(1)
	}

	fmt.Printf("AppName: %s\nVersion: %s\n", cfg.AppName, cfg.Version)

	for _, user := range cfg.Users {
		fmt.Println(user)
	}
}
