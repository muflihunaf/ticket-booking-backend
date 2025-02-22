package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muflihunaf/ticket-booking-v1/models"
)

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/validate", handler.ValidateTicket)
}

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	tickets, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tickets fetched successfully",
		"data":    tickets,
		"status":  "success",
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	ticketId, err := strconv.Atoi(ctx.Params("ticketId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ticket ID",
			"status":  "error",
		})
	}

	ticket, err := h.repository.GetOne(context, uint(ticketId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ticket fetched successfully",
		"data":    ticket,
		"status":  "success",
	})
}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	ticket := &models.Ticket{}
	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ticket data",
			"status":  "error",
		})
	}

	createdTicket, err := h.repository.CreateOne(context, ticket)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ticket created successfully",
		"data":    createdTicket,
		"status":  "success",
	})
}

func (h *TicketHandler) ValidateTicket(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	validateBody := &models.ValidateTicket{}
	if err := ctx.BodyParser(validateBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid validate body",
			"status":  "error",
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	updatedTicket, err := h.repository.UpdateOne(context, validateBody.TicketID, validateData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Ticket validated successfully",
		"data":    updatedTicket,
	})
}
