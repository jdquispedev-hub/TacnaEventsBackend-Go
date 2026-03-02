package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tacna-events-backend/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	db *pgxpool.Pool
}

func NewUserController(db *pgxpool.Pool) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	rows, err := uc.db.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error"})
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}