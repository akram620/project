package repository

import (
	"github.com/akram620/alif/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Events interface {
	Test() *models.Response
}

type EventsRepository struct {
	pool *pgxpool.Pool
}

func NewEventsRepository(pool *pgxpool.Pool) *EventsRepository {
	return &EventsRepository{pool}
}

func (r *EventsRepository) Test() *models.Response {
	return &models.Response{Code: 200, Message: "OK", Payload: time.Now()}
}
