package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"jelastic-golang-hello/internal/application"
	"jelastic-golang-hello/internal/domain"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	user, err := h.userService.CreateUser(c.Context(), req.Name, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidUserData) {
			return c.Status(400).JSON(fiber.Map{"error": "Name and email are required"})
		}
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return c.Status(409).JSON(fiber.Map{"error": "User with this email already exists"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.userService.GetUser(c.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	user, err := h.userService.UpdateUser(c.Context(), uint(id), req.Name, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		if errors.Is(err, domain.ErrInvalidUserData) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid user data"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.userService.DeleteUser(c.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}