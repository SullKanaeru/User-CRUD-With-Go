package main

import (
	"user_crud/internal/config"
	"user_crud/internal/handler"
	"user_crud/internal/repository"
	"user_crud/internal/service"
	"user_crud/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := database.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, userRepo)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api/users")
	api.Post("/", userHandler.Create)
	api.Get("/", userHandler.GetAll)
	api.Put("/:id", userHandler.Update)
	api.Delete("/:id", userHandler.Delete)

	port := config.GetEnv("PORT", "3000")
	app.Listen(":" + port)
}