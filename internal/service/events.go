package service

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/repository"
)

type Events interface {
	Test() *models.Response
}

type EventsService struct {
	eventsRepository repository.Events
}

func NewEventsService(chatRepository repository.Events) *EventsService {
	return &EventsService{chatRepository}
}

func (s *EventsService) Test() *models.Response {
	return s.eventsRepository.Test()
}
