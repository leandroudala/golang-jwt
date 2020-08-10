package repository

import "github.com/leandroudala/golang_jwt/api/models"

// UserRepository ...
type UserRepository interface {
	Save(models.User) (models.User, error)
	All() ([]models.User, error)
	FindByID(uint32) (models.User, int, error)
	// Update(uint32) (models.User, error)
	// Delete(uint32) (int64, error)
}
