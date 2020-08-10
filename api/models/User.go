package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leandroudala/golang_jwt/api/security"
)

// User db model
type User struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Email        string    `gorm:"size:255;unique;not null" json:"email"`
	RegisteredOn time.Time `gorm:"default:current_timestamp();not null" json:"registered_on"`
	Admin        bool      `gorm:"not null" json:"admin,omitempty"`
	PublicID     string    `gorm:"size:36;not null;unique" json:"public_id"`
	Username     string    `gorm:"size:50;not null;unique" json:"username"`
	PasswordHash string    `gorm:"size:100;not null" json:"password,omitempty"`
}

// BeforeSave prepares user before save
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash((u.PasswordHash))
	if err != nil {
		return err
	}

	u.Admin = false
	u.PasswordHash = string(hashedPassword)
	u.PublicID = uuid.New().String()
	u.RegisteredOn = time.Now().UTC()

	return nil
}

// AfterSave cleans password to avoid return after insert
func (u *User) AfterSave() error {
	var boolNil bool

	u.ID = 0
	u.PasswordHash = ""
	u.Admin = boolNil

	return nil
}

// TableName returns the table name...
func (u *User) TableName() string {
	return "user"
}
