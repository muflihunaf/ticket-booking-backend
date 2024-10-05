package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflihunaf/ticket-booking-v1/handlers"
	"github.com/muflihunaf/ticket-booking-v1/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "ticketBooking ",
		ServerHeader: "fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
