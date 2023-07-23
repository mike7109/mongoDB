package main

import (
	"github.com/gofiber/fiber/v2"
	"mongoBD/internal/repositories"
	"mongoBD/internal/service"
	"mongoBD/pkg"
)

func main() {

	dbClient, err := pkg.InitMongo()
	if err != nil {
		panic(err)
	}

	usersRepositories := repositories.NewUsersRepositories(dbClient)
	usersService := service.NewUsersService(usersRepositories)

	app := fiber.New()

	app.Post("/:user_id/:name", usersService.CreateUser)
	app.Get("/:user_id", usersService.GetUser)
	app.Delete("/:user_id", usersService.DeleteUser)

	err = app.Listen(":3002")
	if err != nil {
		panic("error Listen")
	}
}
