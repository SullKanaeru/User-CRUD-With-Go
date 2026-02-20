package handler

import (
	"user_crud/internal/model"
	"user_crud/internal/repository"
	"user_crud/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *service.UserService
	Repo    *repository.UserRepository
}

func NewUserHandler(service *service.UserService, repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Service: service, Repo: repo}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	if err := h.Service.CreateUser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "User berhasil dibuat", "data": user})
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	role := c.Query("role") 

	var users []model.User
	var err error

	if role != "" {
		users, err = h.Repo.FindByRole(role)
	} else {
		users, err = h.Repo.FindAll()
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	
	return c.JSON(fiber.Map{"data": users})
}

func (h *UserHandler) GetByRole(c *fiber.Ctx) error {
	role := c.Params("role")
	users, err := h.Repo.FindByRole(role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(fiber.Map{"data": users})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.Repo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	h.Repo.Update(user)
	return c.JSON(fiber.Map{"message": "User berhasil diupdate", "data": user})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus user"})
	}
	return c.JSON(fiber.Map{"message": "User berhasil dihapus"})
}