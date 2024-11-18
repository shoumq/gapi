package services

import (
	"gapi/pkg/models"
	"gapi/pkg/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *UserRequest) (*models.User, error) {
	user := models.User{Name: req.Name, Email: req.Email, Password: req.Password}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var userPointers []*models.User
	for i := range users {
		userPointers = append(userPointers, &users[i])
	}

	return userPointers, nil
}

func (s *UserService) DeleteUserById(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUserById(id int, req *UserRequest) (*models.User, error) {
	user := models.User{ID: id, Name: req.Name, Email: req.Email, Password: req.Password}

	updatedUser, err := s.repo.Update(id, user)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (s *UserService) AddAdmin(id int) error {
	err := s.repo.AddAdmin(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DelAdmin(id int) error {
	err := s.repo.DelAdmin(id)
	if err != nil {
		return err
	}
	return nil
}
