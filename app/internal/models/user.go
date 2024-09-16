package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	IdentityNumber string    `gorm:"unique" json:"identity_number"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Addresses      []string  `gorm:"type:text[]" json:"addresses"`
	PhoneNumbers   []string  `gorm:"type:text[]" json:"phone_numbers"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}
