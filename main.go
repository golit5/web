package main

import (
	"encoding/json"
	"fmt"
	"myapp/app"
	_ "myapp/docs"
	_ "myapp/internal/controller"
	_ "myapp/internal/model"
	_ "myapp/internal/repository"
	_ "myapp/internal/service"
	"os"
)

// @Description Конфигурация подключения к базе данных
type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
	Sslmode  string `json:"sslmode"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла конфигурации:", err)
		return
	}
	defer file.Close()

	var config DBConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.Dbname, config.Port, config.Sslmode)

	app.StartApp(dsn)
}
