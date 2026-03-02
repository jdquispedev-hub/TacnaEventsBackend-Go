package controllers

import (
	"context"
	"net/http"
	"tacna-events-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventController struct {
	db *pgxpool.Pool
}

func NewEventController(db *pgxpool.Pool) *EventController {
	return &EventController{db: db}
}

func (ec *EventController) GetEvents(c *gin.Context) {
	rows, err := ec.db.Query(context.Background(), "SELECT id, title, description, datetime, location, price, image_url, category, priority FROM events")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error", "details": err.Error()})
		return
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event

		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.DateTime, &event.Location, &event.Price, &event.Image_URL, &event.Category, &event.Priority)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error", "details": err.Error()})
			return
		}

		events = append(events, event)
	}

	c.JSON(http.StatusOK, events)
}
