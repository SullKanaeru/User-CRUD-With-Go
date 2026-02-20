package main

import (
	"user_crud/internal/config" // SESUAIKAN NAMA MODULE
	"user_crud/internal/handler"
	"user_crud/internal/repository"
	"user_crud/internal/service"
	"user_crud/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Setup Database & Jalankan Migrasi Otomatis
	db := database.ConnectDB()

	// 2. Setup Layer
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, userRepo)

	// 3. Setup Fiber App
	app := fiber.New()
	app.Use(logger.New())

	// 4. Setup Routes
	api := app.Group("/api/users")
	api.Post("/", userHandler.Create)
	api.Get("/", userHandler.GetAll)
	api.Put("/:id", userHandler.Update)
	api.Delete("/:id", userHandler.Delete)

	// 5. Start Server
	port := config.GetEnv("PORT", "3000")
	app.Listen(":" + port)
}