package model

import (
	"gorm.io/gorm"
	"time"
)

type URL struct {
	gorm.Model
	ID              int64     `gorm:"primary_key;auto_increment" json:"id"`
	Original_url    string    `gorm:"not null" json:"original_url"`
	Shortened_url   string    `gorm:"not null" json:"shortened_url"`
	Expired_date    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"time"`
}
