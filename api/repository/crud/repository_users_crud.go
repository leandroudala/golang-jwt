package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/leandroudala/golang_jwt/api/models"
	"github.com/leandroudala/golang_jwt/api/utils/channels"
)

// RepositoryUsersCRUD struct
type RepositoryUsersCRUD struct {
	db *gorm.DB
}

// NewRepositoryUsersCRUD ...
func NewRepositoryUsersCRUD(db *gorm.DB) *RepositoryUsersCRUD {
	return &RepositoryUsersCRUD{db}
}

// Save a new user
func (r *RepositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	return models.User{}, err
}
