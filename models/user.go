package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string `json:"username" valid:"required"`
	Email     string `json:"email" valid:"email,required"`
	Password  string `json:"password" valid:"required,min=6"`
	Photos    []Photo
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate() (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now()
	return
}