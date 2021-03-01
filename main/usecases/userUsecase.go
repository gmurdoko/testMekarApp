package usecases

import "testMekarApp/main/models"

//UserUsecase app
type UserUsecase interface {
	GetUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	PostUser(*models.User) error
	PutUser(*models.User) error
	DeleteUser(id int) error
}
