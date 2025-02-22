package repositories

import (
	"context"

	"github.com/muflihunaf/ticket-booking-v1/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	event := []*models.Event{}

	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).First(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}
func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Model(&models.Event{}).Create(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}
func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).Updates(updateData)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}
func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).Delete(&models.Event{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *EventRepository) FindByFilter(ctx context.Context, filters map[string]interface{}) ([]*models.Event, error) {
	var events []*models.Event
	res := r.db.Model(&models.Event{}).Where(filters).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}
	return events, nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
