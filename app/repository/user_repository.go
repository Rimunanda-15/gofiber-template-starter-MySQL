package repository

import (
	"database/sql"
	models "template-starter-mysql/app/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{DB: db}
}

func (user *UserRepository) GetAllUsers() ([]models.User, error){
	rows, err := user.DB.Query("SELECT * FROM users")

	if err != nil{
		return nil,err
	}
	
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}