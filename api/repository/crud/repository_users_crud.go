package crud

import (
	"errors"
	"strings"

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

// Update user
func (r *RepositoryUsersCRUD) Update(user models.User) (models.User, int, error) {
	var err error
	var status int = 200

	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Where("public_id = ?", &user.PublicID).Update(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, status, nil
	}
	if strings.Contains(err.Error(), "Error 1062") {
		status = 409
		err = errors.New("Data conflict")
	} else if gorm.IsRecordNotFoundError(err) {
		status = 404
	} else {
		status = 500
	}

	return models.User{}, status, err
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
func (r *RepositoryUsersCRUD) FindByID(publicID string) (models.User, int, error) {
	var err error
	var status = 200

	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(100).Where("public_id = ?", publicID).First(&user).Error
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

// Delete a user
func (r *RepositoryUsersCRUD) Delete(publicID string) (int, error) {
	var err error
	var status = 200

	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Where("public_id = ?", publicID).Delete(&models.User{}).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return status, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		status = 404
	} else {
		status = 500
	}
	return status, err
}
