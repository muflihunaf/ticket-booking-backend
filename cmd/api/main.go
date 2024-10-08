package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflihunaf/ticket-booking-v1/config"
	"github.com/muflihunaf/ticket-booking-v1/db"
	"github.com/muflihunaf/ticket-booking-v1/handlers"
	"github.com/muflihunaf/ticket-booking-v1/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "ticketBooking ",
		ServerHeader: "fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":" + envConfig.ServerPort)
}
