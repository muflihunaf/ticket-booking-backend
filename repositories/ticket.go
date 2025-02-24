package repositories

import (
	"context"
	"errors"

	"github.com/muflihunaf/ticket-booking-v1/models"
	"gorm.io/gorm"
)

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	ticket := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Preload("Event").Where("id = ?", ticketId).First(&ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	// Check if event exists
	var exists bool
	if err := r.db.Model(&models.Event{}).Select("count(*) > 0").Where("id = ?", ticket.EventID).Find(&exists).Error; err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("event not found")
	}

	// Create ticket and preload event in a single step
	if err := r.db.Create(ticket).Preload("Event").First(ticket, ticket.ID).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	// Fetch the ticket and preload event
	if err := r.db.Preload("Event").First(ticket, ticketId).Error; err != nil {
		return nil, errors.New("ticket not found")
	}

	// Apply updates
	if err := r.db.Model(ticket).Updates(updateData).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) DeleteOne(ctx context.Context, ticketId uint) error {
	res := r.db.Model(&models.Ticket{}).Where("id = ?", ticketId).Delete(&models.Ticket{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
