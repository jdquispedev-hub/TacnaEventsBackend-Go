package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tacna-events-backend/models"
    "tacna-events-backend/db"	
	"context"
)

// func GetUsers(c *gin.Context) {

//     users := []models.User{}

//     c.JSON(http.StatusOK, users)
// }
func GetUsers(c *gin.Context) {

	conn, err := db.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error"})
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}