package repositories

import "gapi/pkg/models"

type UserRepositoryInterface interface {
	Create(user models.User) (models.User, error)
	GetByID(id int) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(id int) error
	Update(user models.User) (models.User, error)
	setAdminStatus(id int, isAdmin bool) error
	AddAdmin(id int) error
	DelAdmin(id int) error
}
