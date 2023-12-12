package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/izumra/serialization/internal/repository"
)

func GetPets(r *repository.File) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pets, err := r.PetsFromFile()
		if err != nil {
			c.Status(fiber.ErrBadGateway.Code)
			return c.JSON(Bad(err))
		}
		return c.JSON(Success(pets))
	}
}

func AddPets(r *repository.File) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pets []repository.Pet
		if err := json.Unmarshal(c.Body(), &pets); err != nil {
			c.Status(fiber.ErrBadRequest.Code)
			return c.JSON(Bad(err))
		}
		r.PetsToFile(pets)
		return c.JSON(Success("Животные записаны в файл"))
	}
}

func DownloadFile(r *repository.File) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Download("peoples.json")
	}
}
