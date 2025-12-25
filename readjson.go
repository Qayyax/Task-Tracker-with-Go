package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBHost       string `json:"database_host"`
	DBPort       int    `json:"database_port"`
	DBUsername   string `json:"database_username"`
	DBPassword   string `json:"database_password"`
	ServerPort   int    `json:"server_port"`
	ServerDebug  bool   `json:"server_debug"`
	SeverTimeout int    `json:"server_timeout"`
}

func readJson() {
	filepath := "config.json"
	var config Config

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}
