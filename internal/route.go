package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izumra/serialization/internal/handler"
	"github.com/izumra/serialization/internal/repository"
)

func SetRoutes(app *fiber.App) {
	r := &repository.File{}
	app.Get("/pets", handler.GetPets(r))
	app.Post("/pets", handler.AddPets(r))
	app.Get("/file", handler.DownloadFile(r))
}
