package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"mongoBD/internal/repositories"
	"mongoBD/internal/user"
	"net/http"
)

var _ UsersRepositories = repositories.UsersRepositories{}

type UsersRepositories interface {
	Create(ctx context.Context, userID int, name string) error
	Get(ctx context.Context, userID int) (user.User, error)
	Delete(ctx context.Context, userID int) error
}

type UsersService struct {
	UsersRepositories UsersRepositories
}

func NewUsersService(UsersRepositories UsersRepositories) UsersService {
	return UsersService{UsersRepositories: UsersRepositories}
}

func (s UsersService) CreateUser(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id", 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}
	name := c.Params("name", "")
	if name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid name"})
	}

	if err = s.UsersRepositories.Create(c.UserContext(), userID, name); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Create failed"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

func (s UsersService) GetUser(c *fiber.Ctx) error {

	var (
		result user.User
	)
	{
	}

	userID, err := c.ParamsInt("user_id", 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	result, err = s.UsersRepositories.Get(c.UserContext(), userID)

	return c.Status(http.StatusOK).JSON(result)
}

func (s UsersService) DeleteUser(c *fiber.Ctx) error {

	userID, err := c.ParamsInt("user_id", 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	if err := s.UsersRepositories.Delete(c.UserContext(), userID); err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
