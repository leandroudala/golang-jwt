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

// TableName returns the table name...
func (u *User) TableName() string {
	return "user"
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

func (u *User) cleanBeforeShow() {
	var boolNil bool
	u.ID = 0
	u.PasswordHash = ""
	u.Admin = boolNil
}

// AfterSave cleans password to avoid return after insert
func (u *User) AfterSave() error {
	u.cleanBeforeShow()
	return nil
}

// AfterFind cleans fields before show
func (u *User) AfterFind() (err error) {
	u.cleanBeforeShow()
	return
}

// BeforeUpdate removes ID and public_id fields, and changes password hash, if informed
func (u *User) BeforeUpdate() (err error) {
	u.ID = 0
	u.PublicID = ""

	if u.PasswordHash == "" {
		return
	}

	hashedPassword, err := security.Hash((u.PasswordHash))
	u.PasswordHash = string(hashedPassword)

	return
}
