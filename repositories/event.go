package repositories

import (
	"context"
	"time"

	"github.com/muflihunaf/ticket-booking-v1/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "2132132132131",
		Name:      "My Fav Name",
		Location:  "Somewhere",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}
func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
