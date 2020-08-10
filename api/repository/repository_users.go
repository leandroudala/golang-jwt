package repository

import "github.com/leandroudala/golang_jwt/api/models"

// UserRepository ...
type UserRepository interface {
	Save(models.User) (models.User, int, error)
	All() ([]models.User, error)
	FindByID(string) (models.User, int, error)
	Update(models.User) (models.User, int, error)
	Delete(string) (int, error)
}
