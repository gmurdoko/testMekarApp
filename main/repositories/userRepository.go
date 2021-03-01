package repositories

import "testMekarApp/main/models"

type UserRepository interface {
	SelectUsers() ([]*models.User, error)
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	SelectUserById(id int) (*models.User, error)
}
