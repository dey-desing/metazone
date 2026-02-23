package services

import (
	"metazone/models"
	"metazone/db"
)

// Lista de usuarios en memoria
var users []models.User

// Crear un nuevo usuario
func CreateUser(user models.User) (*models.User, error) {

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := db.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return &models.User{
		ID:    int(id),
		Name:  user.Name,
		Email: user.Email,
		Password: user.Password,
	}, nil
}

func GetUsers() []models.User {

	rows, _ := db.DB.Query("SELECT id, name, email FROM users")
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	return users
}