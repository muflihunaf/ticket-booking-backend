package repositories

import (
	"context"

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

	res := r.db.Model(&models.Ticket{}).Preload("Event").Where("id = ?", ticketId).First(&ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(&models.Ticket{}).Create(ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Where("id = ?", ticketId).Updates(updateData)
	if res.Error != nil {
		return nil, res.Error
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
