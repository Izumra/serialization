package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/izumra/serialization/internal"
)

func main() {
	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.Config{}))
	app.Static("/", "./static")
	internal.SetRoutes(app)
	fmt.Println("Увидел, что нужно было в xml сделать сериализацию, было уже поздно :(")

	fmt.Println("Можно переходить по ссылке http://localhost:6670")
	fmt.Println("")
	fmt.Println("/file скачает файл json")
	fiberlog.Error(app.Listen(":6670"))
}
