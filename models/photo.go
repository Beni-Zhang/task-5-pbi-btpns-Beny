package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	Title     string `json:"title" valid:"required"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url" valid:"required,url"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) BeforeCreate() (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

func (p *Photo) BeforeUpdate() (err error) {
	p.UpdatedAt = time.Now()
	return
}