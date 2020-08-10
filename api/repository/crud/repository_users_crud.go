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

// All list all users
func (r *RepositoryUsersCRUD) All() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

// FindByID returns a user
func (r *RepositoryUsersCRUD) FindByID(id uint32) (models.User, int, error) {
	var err error
	var status = 200

	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(100).Where("id = ?", id).First(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, status, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		status = 404
	} else {
		status = 500
	}
	return models.User{}, status, err
}
