package handler

import "github.com/gofiber/fiber/v2"

func Success(data interface{}) fiber.Map {
	return fiber.Map{
		"data":  data,
		"error": nil,
	}
}
func Bad(err error) fiber.Map {
	return fiber.Map{
		"data":  nil,
		"error": err.Error(),
	}
}
