package services

import (
	"context"
	"database/sql"
	"tacna-events-backend/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventService struct {
	db *pgxpool.Pool
}

func NewEventService(db *pgxpool.Pool) *EventService {
	return &EventService{db: db}
}

func (es *EventService) GetAllEvents() ([]models.Event, error) {
	query := `
		SELECT id, title, description, datetime, location, price, 
		       image_url, category_id, priority, created_at, updated_at 
		FROM events 
		ORDER BY datetime ASC`

	rows, err := es.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var categoryID sql.NullInt64
		var imageURL sql.NullString
		var location sql.NullString
		var price sql.NullFloat64

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.DateTime,
			&location,
			&price,
			&imageURL,
			&categoryID,
			&event.Priority,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			continue
		}

		// Handle nullable fields
		if location.Valid {
			event.Location = location.String
		}
		if price.Valid {
			event.Price = price.Float64
		}
		if imageURL.Valid {
			event.Image_URL = imageURL.String
		}
		if categoryID.Valid {
			event.CategoryID = int(categoryID.Int64)
		}

		events = append(events, event)
	}

	return events, nil
}

func (es *EventService) GetEventByID(id int) (*models.Event, error) {
	query := `
		SELECT id, title, description, datetime, location, price, 
		       image_url, category_id, priority, created_at, updated_at 
		FROM events 
		WHERE id = $1`

	var event models.Event
	var categoryID sql.NullInt64
	var imageURL sql.NullString
	var location sql.NullString
	var price sql.NullFloat64

	err := es.db.QueryRow(context.Background(), query, id).Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.DateTime,
		&location,
		&price,
		&imageURL,
		&categoryID,
		&event.Priority,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Handle nullable fields
	if location.Valid {
		event.Location = location.String
	}
	if price.Valid {
		event.Price = price.Float64
	}
	if imageURL.Valid {
		event.Image_URL = imageURL.String
	}
	if categoryID.Valid {
		event.CategoryID = int(categoryID.Int64)
	}

	return &event, nil
}

func (es *EventService) CreateEvent(event *models.Event) error {
	query := `
		INSERT INTO events (title, description, datetime, location, price, image_url, category_id, priority)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at`

	var location, imageURL sql.NullString
	var price sql.NullFloat64
	var categoryID sql.NullInt64

	// Convert to nullable types
	if event.Location != "" {
		location.String = event.Location
		location.Valid = true
	}
	if event.Price > 0 {
		price.Float64 = event.Price
		price.Valid = true
	}
	if event.Image_URL != "" {
		imageURL.String = event.Image_URL
		imageURL.Valid = true
	}
	if event.CategoryID > 0 {
		categoryID.Int64 = int64(event.CategoryID)
		categoryID.Valid = true
	}

	err := es.db.QueryRow(context.Background(), query,
		event.Title,
		event.Description,
		event.DateTime,
		location,
		price,
		imageURL,
		categoryID,
		event.Priority,
	).Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt)

	return err
}

func (es *EventService) UpdateEvent(id int, event *models.Event) error {
	query := `
		UPDATE events 
		SET title = $1, description = $2, datetime = $3, location = $4, 
		    price = $5, image_url = $6, category_id = $7, priority = $8,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $9
		RETURNING updated_at`

	var location, imageURL sql.NullString
	var price sql.NullFloat64
	var categoryID sql.NullInt64

	// Convert to nullable types
	if event.Location != "" {
		location.String = event.Location
		location.Valid = true
	}
	if event.Price > 0 {
		price.Float64 = event.Price
		price.Valid = true
	}
	if event.Image_URL != "" {
		imageURL.String = event.Image_URL
		imageURL.Valid = true
	}
	if event.CategoryID > 0 {
		categoryID.Int64 = int64(event.CategoryID)
		categoryID.Valid = true
	}

	err := es.db.QueryRow(context.Background(), query,
		event.Title,
		event.Description,
		event.DateTime,
		location,
		price,
		imageURL,
		categoryID,
		event.Priority,
		id,
	).Scan(&event.UpdatedAt)

	return err
}

func (es *EventService) DeleteEvent(id int) error {
	query := `DELETE FROM events WHERE id = $1`

	_, err := es.db.Exec(context.Background(), query, id)
	return err
}
