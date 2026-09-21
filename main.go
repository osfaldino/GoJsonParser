package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct for json data
type Config struct {
	AppName string   `json:"appName"`
	Version string   `json:"version"`
	Debug   bool     `json:"debug"`
	Users   []string `json:"users"`
}

// func to read file
func (cfg *Config) ParseConfig(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close() // close file at the end

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg) // parse json into struct
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var cfg Config
	err := cfg.ParseConfig("cfg.json") // read config
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	// print result
	fmt.Printf("AppName: %s\nVersion: %s\n", cfg.AppName, cfg.Version)

	for _, user := range cfg.Users {
		fmt.Println(user)
	}
}
