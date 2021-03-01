package usecases

import (
	"testMekarApp/main/models"
	"testMekarApp/main/repositories"
)

// UserUsecaseImpl app
type UserUsecaseImpl struct {
	UserRepository repositories.UserRepository
}

// GetUserById app
func (s UserUsecaseImpl) GetUserById(id int) (*models.User, error) {
	user, err := s.UserRepository.SelectUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserUsecaseImpl) GetUsers() ([]*models.User, error) {
	users, err := s.UserRepository.SelectUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s UserUsecaseImpl) PostUser(user *models.User) error {
	err := s.UserRepository.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserUsecaseImpl) PutUser(user *models.User) error {
	err := s.UserRepository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserUsecaseImpl) DeleteUser(id int) error {
	err := s.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

//InitUserusecaseImpl app
func InitUserusecaseImpl(userRepository repositories.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepository}
}
