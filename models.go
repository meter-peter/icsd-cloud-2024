package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	City    string `json:"city"`
	Country string `json:"country"`
}
