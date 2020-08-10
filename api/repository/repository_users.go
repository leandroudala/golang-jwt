package repository

import "github.com/leandroudala/golang_jwt/api/models"

// UserRepository ...
type UserRepository interface {
	Save(models.User) (models.User, error)
	// FindAll() ([]models.User, error)
	// FindById(uint32) (models.User, error)
	// Update(uint32) (models.User, error)
	// Delete(uint32) (int64, error)
}
