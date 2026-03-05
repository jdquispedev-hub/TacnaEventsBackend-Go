package services

import (
	"tacna-events-backend/models"

	"gorm.io/gorm"
)

type EventService struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{db: db}
}

func (es *EventService) GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	result := es.db.Find(&events)
	return events, result.Error
}

func (es *EventService) GetEventByID(id uint) (*models.Event, error) {
	var event models.Event
	result := es.db.First(&event, id)
	return &event, result.Error
}

func (es *EventService) CreateEvent(event *models.Event) error {
	result := es.db.Create(event)
	return result.Error
}

func (es *EventService) UpdateEvent(id uint, event *models.Event) error {
	result := es.db.Model(&models.Event{}).Where("id = ?", id).Updates(event)
	return result.Error
}

func (es *EventService) DeleteEvent(id uint) error {
	result := es.db.Delete(&models.Event{}, id)
	return result.Error
}
