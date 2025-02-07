package main

import (
	"log"
	"myapp/internal/controller"
	"myapp/internal/model"
	"myapp/internal/repository"
	"myapp/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключаемся к базе данных
	dsn := "host=localhost user=myuser password=mypassword dbname=myapp_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Миграция базы данных
	db.AutoMigrate(&model.Link{})

	// Настройка репозитория, сервиса и контроллера
	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkController := controller.NewLinkController(linkService)

	// Настройка Gin роутера
	r := gin.Default()

	// Настройка CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Разрешить все источники или замените на конкретные URL
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},        // Разрешаем методы, включая OPTIONS для preflight-запросов
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},                          // Заголовки, которые могут быть видны клиенту
		AllowCredentials: true,                                                // Разрешаем cookies
		MaxAge:           12 * 3600,                                           // Время кэширования preflight-запросов
	}))

	// Регистрация API маршрутов
	r.POST("/links", linkController.CreateLink)
	r.GET("/links", linkController.GetAllLinks)
	r.GET("/links/:id", linkController.GetLinkByID)
	r.PUT("/links/:id", linkController.UpdateLink)
	r.DELETE("/links/:id", linkController.DeleteLink)

	// Запуск сервера
	r.Run(":8080")
}
