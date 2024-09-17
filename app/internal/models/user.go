package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}

	var arr []string
	err := json.Unmarshal(bytes, &arr)
	*a = StringArray(arr)
	return err
}

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return json.Marshal(a)
}

type User struct {
	ID             uuid.UUID   `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt      time.Time   `json:"created_at"`
	IdentityNumber string      `gorm:"unique" json:"identity_number"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Gender         string      `json:"gender"`
	DateOfBirth    time.Time   `json:"date_of_birth"`
	Addresses      StringArray `gorm:"type:jsonb" json:"addresses"`
	PhoneNumbers   StringArray `gorm:"type:jsonb" json:"phone_numbers"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}
